package main

import (
    "log"
    "github.com/Kitagawa19/GO_prefecture/database/db"
    "github.com/Kitagawa19/GO_prefecture/handlers/address"
    "github.com/gin-gonic/gin"
)
func main() {
    r :=gin.Default()
    r.LoadHTMLGlob("templates/*")
    dn,err := database.InitDB("new.db")
    if err != nil {
        log.Fatal("Failed to initialize database: %v",err)
    }
    defer db.Close()

    if err := database.ImportCSV("Data/utf_ken_all.csv"); err != nil {
        lo9g.Fatal("Failed to import CSV: %v",err)
    }

    r.GET("/".handlers.IndexHandler)
    r.POST("/search",handlers.SearchHandler)
    r.RUN(":8080")
}