package main

import (
	"Go_prefecture/database"
	"Go_prefecture/handlers"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	db, err := database.InitDB("new.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 最初のテーブルを作成
	fmt.Println("First table created successfully")

	// CSVデータをaddressesテーブルにインポート
	err = database.ImportCSV("Data/utf_ken_all.csv") // CSVファイルのパスを指定
	if err != nil {
		fmt.Printf("Error importing CSV: %v\n", err)
		return
	}
	fmt.Println("CSV data imported successfully")

	// addresses テーブルにデータが存在するか確認
	rows, err := db.Query("SELECT COUNT(*) FROM addresses")
	if err != nil {
		fmt.Printf("Error querying addresses table: %v\n", err)
		return
	}
	var count int
	if rows.Next() {
		rows.Scan(&count)
	}
	fmt.Printf("Addresses table contains %d records\n", count)
	rows.Close()

	if count == 0 {
		fmt.Println("Addresses table is empty. No data to normalize.")
		return
	}

	// NormalizeTableを呼び出して、2つ目のテーブルを作成
	err = database.NormalizeTable()
	if err != nil {
		fmt.Printf("Error occurred during normalization: %v\n", err)
		return
	}

	fmt.Println("Second table created successfully")

	// 2つ目のテーブルにデータが存在するか確認
	rows, err = db.Query("SELECT COUNT(*) FROM normalized_utf_ken_all")
	if err != nil {
		fmt.Printf("Error querying normalized_utf_ken_all table: %v\n", err)
		return
	}
	if rows.Next() {
		rows.Scan(&count)
	}
	fmt.Printf("normalized_utf_ken_all table contains %d records\n", count)
	rows.Close()

	router.GET("/", handlers.HomeHandler)
	router.GET("/prefectures", handlers.PrefectureHandler)  //都道府県リスト
	router.GET("/cities", handlers.PrefecturetocityHandler) //市区町村検索ページ
	router.POST("/citiesresult", handlers.CityHandler)      //市区町村検索結果のページ
	router.GET("/postcode", handlers.PostalHandler)         //郵便番号の検索ページ
	router.POST("/addressresult", handlers.AddressHandler)  //郵便番号から住所の結果を表示するページ
	router.POST("/postsearch", handlers.PostSearchHandler)  //郵便番号から住所の結果を表示するページ
	router.Run(":8080")
}
