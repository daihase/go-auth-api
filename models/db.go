package models

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "os"
)

var DB *sql.DB

func InitDB() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }
}
