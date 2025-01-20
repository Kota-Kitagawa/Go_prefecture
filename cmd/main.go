package main

import (
    "Go_prefecture/internal/database"
    "Go_prefecture/handlers"
    "github.com/gin-gonic/gin"
    "fmt"
    "html/template"
    "os"
)

func main() {
    database_url :=os.Getenv("DATABASE_URL")
    router := gin.Default()
    router.SetFuncMap(template.FuncMap{
        "add": func(a, b int) int { return a + b },
        "sub": func(a, b int) int { return a - b },
    })
    router.LoadHTMLGlob("src/templates/*")
    db, err := database.InitDB(database_url)
    if err != nil {
        panic(err)
    }
    defer db.Close()
    fmt.Println("First table created successfully")

    // CSVデータをaddressesテーブルにインポート
    err = database.ImportCSV("./Data/utf_ken_all.csv") 
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
    router.GET("/prefectures", handlers.PretoCityHandler) //都道府県リストページ
    router.GET("/cities", handlers.PrefListHTMLHandler)   //市区町村検索ページ
    router.GET("/citiesresult", handlers.CitiesHandler)        //市区町村検索結果のページ
    router.GET("/postcode", handlers.PostalHandler)           //郵便番号の検索ページ
    router.GET("/addressresult", handlers.AddressHandler)    //郵便番号から住所の結果を表示するページ
    router.GET("/postsearch", handlers.AddressSearchHandler)  //住所から郵便番号を検索するページ
    router.GET("/postresult",handlers.PostSearchHandler)     //郵便番号結果を表示するページ
    router.Run(":8080")
}
