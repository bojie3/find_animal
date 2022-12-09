package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB = setupDatabase()

func setupDatabase() *sql.DB {
	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "yyzkm61200"
		dbname   = "animal"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", //need to change when uploading
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	setTables(db)

	fmt.Println("Databse connected and set up successfully!")
	return db
}
