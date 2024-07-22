package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func main() {
	var err error
	if err != nil {
		log.Fatal(err)
	}
	DbConnection, _ := sql.Open("sqlite3", "new.db") //接続開始（example.sqlに保存する）
	defer DbConnection.Close()                       //最後は確実にクローズする。

	//この下に、CREATE文　SELECT文　INSERT文 UPDATE文　DELETE文を記載する
	readUserSQL := `SELECT *  FROM  utf_ken_all LIMIT 5`
	rows, err := DbConnection.Query(readUserSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

}
