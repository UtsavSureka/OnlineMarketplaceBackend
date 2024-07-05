package DBConnection

import (
	"Ecomm/config"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	User_DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME))

	if err != nil {
		return nil, errors.New("Unable to connect to db" + err.Error())
	}

	//Ping the database to ensure successful connection :

	if err = User_DB.Ping(); err != nil {
		log.Fatal(err)
	}

	return User_DB, nil

}
