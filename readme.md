# Knjižnica
Knjižnica je preprost api za izposojo in vračanje knjig napisan v GoLang jeziku.
- database.go je datoteka, ki vsebuje metode za poizvedbe v bazi
- endpoints.go je impementacija REST endpointov
# Zahteve 
- PostgreSQL 10
- GoLang 1.10
# Knjižnice
- database/sql
- github.com/gorilla/mux
# Priprava
- Najprej namestimo PostgreSQL 
- Ustvarimo Uporabnika viberate z geslom viberate
- V bazi z imenom postgres ustvarimo tabele iz datoteke baza.sql
- Namestimo GoLang 
- Namestimo database/sql knjižnico (ukaz "go get -u github.com/lib/pq")
- Namestimo gorilla-mux ("go get -u github.com/gorilla/mux")
# GET endpointi
- http://localhost:8080/users nam vrne seznam uporabnikov skupaj s knjigami ki jih imajo izposojene
- http://localhost:8080/users/{id_uporabnika} nam vrne podatke o uporabniku z id-jem id_uporabnika
- http://localhost:8080/books nam vrne podatke o knjigah ki so na voljo

# POST endpointi
- http://localhost:8080/adduser doda uporabnika. V telesu moramo poslati json z atributi "Ime" in "Priimek"
- primer: {"Ime":"Marie","Priimek":"Curie"}
- http://localhost:8080/borrow nam sposodi knjigo. V telesu moramo podati id knjige, ki si jo želimo izposoditi ("BookID") in id uporabnika, ki si knjigo sposodi ("UserID")
- primer: {"BookID": 2,"UserID": 1}
- http://localhost:8080/return vrne knjigo. V telseu podamo isti json kot pri /borrow