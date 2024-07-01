package DBConnection

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	User_DB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/MarketPlace")
	if err != nil {
		return nil, errors.New("Unable to connect to db" + err.Error())
	}

	//Ping the database to ensure successful connection :

	if err = User_DB.Ping(); err != nil {
		log.Fatal("Failed to ping to database")
	}

	return User_DB, nil

}
