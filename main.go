package main

import (
    "log"
    "Go_prefecture/database"
    "Go_prefecture/handlers"
    "github.com/gin-gonic/gin"
)
func main() {
    router :=gin.Default()
    router.LoadHTMLGlob("templates/*")
    db,err := database.InitDB("new.db")
    if err != nil {
        log.Fatal("Failed to initialize database: %v",err)
    }
    defer db.Close()

    if err := database.ImportCSV("Data/utf_ken_all.csv"); err != nil {
        log.Fatal("Failed to import CSV: %v",err)
    }

    router.GET("/", handlers.HomeHandler)
	router.GET("/prefectures", handlers.PrefectureHandler)
	router.GET("/cities", handlers.CityHandler)
	router.GET("/address", handlers.AddressHandler)
    router.Run(":8080")
}