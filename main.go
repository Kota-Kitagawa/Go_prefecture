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
	router.GET("/prefectures", handlers.PrefectureHandler) //都道府県リスト
	router.GET("/cities", handlers.PrefecturetocityHandler) //市区町村検索ページ
    router.POST("/citiesresult", handlers.CityHandler) //市区町村検索結果のページ
	router.GET("/postcode", handlers.PostalHandler) //郵便番号の検索ページ
    router.POST("/addressresult", handlers.AddressHandler) //郵便番号から住所の結果を表示するページ
    router.Run(":8080")
}