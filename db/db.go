package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(user, pass, name string) *sql.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", user, pass, name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	log.Println("Database connected 🚀")
	return db
}
