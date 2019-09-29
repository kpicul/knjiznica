package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "viberate"
	password = "viberate"
	dbname   = "postgres"
)

type Knjiga struct {
	id       int
	naziv    string
	kolicina int
}

type Uporabnik struct {
	id      int
	ime     string
	priimek string
}

type Lastnistvo struct {
	idKnjiga    int
	idUporabnik int
}

func connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func getUsers() []Uporabnik {
	db := connect()
	defer db.Close()
	sqlSt := "SELECT * FROM uporabnik;"
	var users []Uporabnik
	rows, err2 := db.Query(sqlSt)
	if err2 != nil {
		panic(err2)
	}
	for rows.Next() {
		var user Uporabnik
		err3 := rows.Scan(&user.id, &user.ime, &user.priimek)
		if err3 != nil {
			panic(err3)
		}
		users = append(users, user)
	}
	return users
}

func getBooks() []Knjiga {
	db := connect()
	defer db.Close()
	sqlSt := "SELECT * FROM knjiga;"
	var books []Knjiga
	rows, err2 := db.Query(sqlSt)
	if err2 != nil {
		panic(err2)
	}
	for rows.Next() {
		var book Knjiga
		err3 := rows.Scan(&book.id, &book.naziv, &book.kolicina)
		if err3 != nil {
			panic(err3)
		}
		books = append(books, book)
	}
	return books
}

func insertUser(ime, priimek string) {
	db := connect()
	defer db.Close()
	sqlSt := "INSERT INTO uporabnik (ime, priimek) VALUES ($1, $2);"
	_, err := db.Exec(sqlSt, ime, priimek)
	if err != nil {
		panic(err)
	}
}

func updateQuant(bookId, quant int) {
	db := connect()
	defer db.Close()
	sqlSt := "UPDATE knjiga SET kolicina=$1 WHERE id = $2;"
	_, err := db.Exec(sqlSt, quant, bookId)
	if err != nil {
		panic(err)
	}
}

func insertBorrow(userId, bookId int) {
	db := connect()
	defer db.Close()
	sqlSt := "INSERT INTO lastnistvo(knjiga_id, uporabnik_id) VALUES ($1, $2);"
	_, err := db.Exec(sqlSt, bookId, userId)
	if err != nil {
		panic(err)
	}
}

func getBook(bookId int) Knjiga {
	db := connect()
	defer db.Close()
	bookChk := "SELECT * FROM knjiga WHERE ID=$1;"
	row := db.QueryRow(bookChk, bookId)
	var book Knjiga
	err := row.Scan(&book.id, &book.naziv, &book.kolicina)
	if err != nil {
		panic(err)
	}
	return book
}

func borrow(userId, bookId int) string {
	book := getBook(bookId)
	if book.kolicina == 0 {
		return "Book isn't available"
	}
	insertBorrow(userId, bookId)
	updateQuant(bookId, book.kolicina-1)
	return "Book borrowed"
}

func returnBook(userId, bookId int) string {
	db := connect()
	defer db.Close()
	sqlSt := "DELETE FROM lastnistvo WHERE knjiga_id = $1 AND uporabnik_id = $2"
	_, err := db.Exec(sqlSt, bookId, userId)
	if err != nil {
		panic(err)
	}
	book := getBook(bookId)
	updateQuant(bookId, book.kolicina+1)
	return "Knjiga vrnjena"
}

func getUser(userId int) Uporabnik {
	db := connect()
	defer db.Close()
	sqlSt := "SELECT * FROM uporabnik WHERE id=$1"
	row := db.QueryRow(sqlSt, userId)
	var user Uporabnik
	err := row.Scan(&user.id, &user.ime, &user.priimek)
	if err != nil {
		panic(err)
	}
	return user
}

func getUserBooks(userId int) []Knjiga {
	db := connect()
	defer db.Close()
	booksResult := make([]Knjiga, 0)
	sqlSt := "SELECT k.id, k.naziv, k.kolicina FROM lastnistvo l, knjiga k, uporabnik u WHERE u.id = l.uporabnik_id AND k.id = l.knjiga_id AND l.uporabnik_id = $1"
	rows, err := db.Query(sqlSt, userId)
	if err != nil {
		panic(err)
	}
	var result Knjiga
	for rows.Next() {
		rows.Scan(&result.id, &result.naziv, &result.kolicina)
		booksResult = append(booksResult, result)
	}
	return booksResult
}

func getAvailableBooks() []Knjiga {
	db := connect()
	defer db.Close()
	sqlSt := "SELECT * FROM knjiga WHERE kolicina > 0;"
	var books []Knjiga
	rows, err2 := db.Query(sqlSt)
	if err2 != nil {
		panic(err2)
	}
	for rows.Next() {
		var book Knjiga
		err3 := rows.Scan(&book.id, &book.naziv, &book.kolicina)
		if err3 != nil {
			panic(err3)
		}
		books = append(books, book)
	}
	return books
}
