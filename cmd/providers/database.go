package providers

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

func ConnectDB() {
	connStr := os.Getenv("DB_URL")
	var err error
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}
}
