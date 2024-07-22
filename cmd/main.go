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
    DbConnection, err = sql.Open("sqlite3", "new.db")
    if err != nil {
        log.Fatal(err)
    }
    defer DbConnection.Close()

    // SQLクエリの実行
    readUserSQL := `SELECT * FROM utf_ken_all LIMIT 5`
    rows, err := DbConnection.Query(readUserSQL)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // 取得したデータを表示
    var columns []string
    columns, err = rows.Columns()
    if err != nil {
        log.Fatal(err)
    }
    
    values := make([]interface{}, len(columns))
    valuePtrs := make([]interface{}, len(columns))
    
    for rows.Next() {
        for i := range columns {
            valuePtrs[i] = &values[i]
        }
        
        if err := rows.Scan(valuePtrs...); err != nil {
            log.Fatal(err)
        }
        
        for i, col := range columns {
            val := values[i]
            var v interface{}
            if b, ok := val.([]byte); ok {
                v = string(b)
            } else {
                v = val
            }
            fmt.Printf("%s: %v\n", col, v)
        }
        fmt.Println("-------------------")
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}
