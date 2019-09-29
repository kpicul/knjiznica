DROP TABLE IF EXISTS lastnistvo;
DROP TABLE IF EXISTS knjiga;
DROP TABLE IF EXISTS uporabnik;
CREATE TABLE knjiga (
    id SERIAL PRIMARY KEY,
    naziv varchar(50) NOT NULL,
    kolicina integer
);

CREATE TABLE uporabnik(
    id SERIAL PRIMARY KEY,
    ime varchar(20),
    priimek varchar(20)
);

CREATE TABLE lastnistvo(
    knjiga_id integer NOT NULL,
    uporabnik_id integer NOT NULL,
    FOREIGN KEY (knjiga_id) REFERENCES knjiga(id),
    FOREIGN KEY (uporabnik_id) REFERENCES uporabnik(id)
);

INSERT INTO uporabnik(ime, priimek) values ('test', 'testni');
INSERT INTO uporabnik(ime, priimek) values ('Josip', 'Broz-Tito');
INSERT INTO uporabnik(ime, priimek) values ('Janez', 'Janša');

INSERT INTO knjiga(naziv, kolicina) values ('Dune', '10');
INSERT INTO knjiga(naziv, kolicina) values ('1984', '2');
INSERT INTO knjiga(naziv, kolicina) values ('Hobit', '3');

INSERT INTO lastnistvo(knjiga_id, uporabnik_id) VALUES (2, 1);
INSERT INTO lastnistvo(knjiga_id, uporabnik_id) VALUES (1, 2);
INSERT INTO lastnistvo(knjiga_id, uporabnik_id) VALUES (2, 2);
INSERT INTO lastnistvo(knjiga_id, uporabnik_id) VALUES (1, 3);
