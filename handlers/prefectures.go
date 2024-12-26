package handlers

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "Go_prefecture/internal/database"
)

func fetchPrefecture() ([]string, error) {
    query := `SELECT DISTINCT field7 FROM addresses`
    rows, err := database.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var prefectures []string
    for rows.Next() {
        var prefecture string
        if err := rows.Scan(&prefecture); err != nil {
            return nil, err
        }
        prefectures = append(prefectures, prefecture)
    }
    return prefectures, nil
}

func PrefectureHTMLHandler(c *gin.Context) {
    prefectures, err := fetchPrefecture()
    if err != nil {
        log.Printf("Failed to fetch prefectures: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch prefectures")
        return
    }
    c.HTML(http.StatusOK, "prefectures.html", gin.H{
        "Prefectures": prefectures,
    })
}

func PrefectureJSONHandler(c *gin.Context) {
    prefectures, err := fetchPrefecture()
    if err != nil {
        log.Printf("Failed to fetch prefectures: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch prefectures"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "Prefectures": prefectures,
    })
}

func PrefListHTMLHandler(c *gin.Context) {
    prefectures, err := fetchPrefecture()
    if err != nil {
        log.Printf("Failed to fetch prefectures: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch prefectures")
        return
    }
    c.HTML(http.StatusOK, "cities.html", gin.H{
        "Prefectures": prefectures,
    })
}