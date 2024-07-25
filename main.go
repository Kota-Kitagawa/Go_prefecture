package main

import (
    "log"
    "GO_prefecture/database"
    "GO_prefecture/handlers"
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

    if err := database.ImportCSV("database/Data/utf_ken_all.csv"); err != nil {
        lo9g.Fatal("Failed to import CSV: %v",err)
    }
    
    r.GET("/".handlers.IndexHandler)
    r.POST("/search",handlers.SearchHandler)
    r.RUN(":8080")
}