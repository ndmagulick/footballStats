package database

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var Db *sql.DB

func InitDB() {

	connString := "Server=localhost\\SQLEXPRESS;Database=FootballStats;Trusted_Connection=True;user id=fsAdmin; password=password;"

	log.Printf("Opening connection to database...")

	db, err := sql.Open("mssql", connString)

	if err != nil {
		log.Fatal("Open connection failed: ", err.Error())
	}

	log.Printf("Connection to database opened!")

	log.Printf("Pinging database...")

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	log.Printf("Database returned ping!")

	Db = db
}
