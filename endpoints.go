package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type userJson struct {
	ID        int          `json:"ID"`
	Ime       string       `json:"Ime"`
	Priimek   string       `json:"Priimek"`
	Sposojeno []knjigaJson `json:"Sposojeno"`
}

type createdUser struct {
	Ime     string `json:"Ime"`
	Priimek string `json: "Priimek`
}

type knjigaJson struct {
	ID    int    `json: "ID"`
	Naziv string `json:"Naziv"`
}

type availableBooks struct {
	ID       int    `json :"ID"`
	Naziv    string `json : "Naziv"`
	Kolicina int    `json : "Kolicina"`
}

type borrowReturn struct {
	BookID int `json : "BookID"`
	UserID int `json : "UserID"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")
}

func getUserJson(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	user := getUser(userId)
	var uJson userJson
	uJson.ID = user.id
	uJson.Ime = user.ime
	uJson.Priimek = user.priimek
	uJson.Sposojeno = make([]knjigaJson, 0)
	bookList := getUserBooks(userId)
	for i := 0; i < len(bookList); i++ {
		var book knjigaJson
		book.ID = bookList[i].id
		book.Naziv = bookList[i].naziv
		uJson.Sposojeno = append(uJson.Sposojeno, book)
	}
	json.NewEncoder(w).Encode(uJson)
}
func getAllUsersJson(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	resultUs := make([]userJson, 0)
	for i := 0; i < len(users); i++ {
		var user userJson
		user.ID = users[i].id
		user.Ime = users[i].ime
		user.Priimek = users[i].priimek
		user.Sposojeno = make([]knjigaJson, 0)
		bookList := getUserBooks(user.ID)
		fmt.Println(bookList)
		for i := 0; i < len(bookList); i++ {
			var book knjigaJson
			book.ID = bookList[i].id
			book.Naziv = bookList[i].naziv
			user.Sposojeno = append(user.Sposojeno, book)
		}
		fmt.Println(user)
		resultUs = append(resultUs, user)
	}
	json.NewEncoder(w).Encode(resultUs)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user createdUser
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong input")
	}
	json.Unmarshal(reqBody, &user)
	insertUser(user.Ime, user.Priimek)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func retAvailableBooks(w http.ResponseWriter, r *http.Request) {
	books := make([]availableBooks, 0)
	bookDb := getAvailableBooks()
	for i := 0; i < len(bookDb); i++ {
		var booka availableBooks
		booka.ID = bookDb[i].id
		booka.Naziv = bookDb[i].naziv
		booka.Kolicina = bookDb[i].kolicina
		books = append(books, booka)
	}
	json.NewEncoder(w).Encode(books)
}

func borrowBook(w http.ResponseWriter, r *http.Request) {
	var bor borrowReturn
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(reqBody, &bor)
	borrow(bor.UserID, bor.BookID)
	bookname := getBook(bor.BookID)
	fmt.Fprintf(w, "Book "+bookname.naziv+" borrowed")
}

func returnBorrowedBook(w http.ResponseWriter, r *http.Request) {
	var bor borrowReturn
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(reqBody, &bor)
	returnBook(bor.UserID, bor.BookID)
	bookname := getBook(bor.BookID)
	fmt.Fprintf(w, "Book "+bookname.naziv+" returned")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/users/{id}", getUserJson).Methods("GET")
	router.HandleFunc("/users", getAllUsersJson).Methods("GET")
	router.HandleFunc("/adduser", createUser).Methods("POST")
	router.HandleFunc("/books", retAvailableBooks).Methods("GET")
	router.HandleFunc("/borrow", borrowBook).Methods("POST")
	router.HandleFunc("/return", returnBorrowedBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
