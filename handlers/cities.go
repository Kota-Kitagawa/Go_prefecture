package handlers

import (
    "log"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "Go_prefecture/internal/database"
)

func fetchCities(prefecture string, limit, offset int) ([]string, error) {
    query := `
        SELECT CASE
            WHEN field9 = '以下に掲載がない場合' THEN field8
            ELSE field8 || field9
            END AS city
        FROM addresses
        WHERE field7 = ?
        LIMIT ? OFFSET ?
    `
    rows, err := database.DB.Query(query, prefecture, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var cities []string
    for rows.Next() {
        var city string
        if err := rows.Scan(&city); err != nil {
            return nil, err
        }
        cities = append(cities, city)
    }
    return cities, err
}

func CitiesHTMLHandler(c *gin.Context) {
    prefecture := c.Query("prefecture")
    pageStr := c.Query("page")
    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        page = 1
    }
    limit := 10
    offset := (page - 1) * limit

    log.Printf("Received prefecture: %s, page: %d", prefecture, page)

    cities, err := fetchCities(prefecture, limit, offset)
    if err != nil {
        log.Printf("Failed to fetch cities: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch cities")
        return
    }

    c.HTML(http.StatusOK, "citiesresult.html", gin.H{
        "Cities": cities,
        "Page":   page,
        "Prefecture": prefecture,
    })
}

func CitiesJSONHandler(c *gin.Context) {
    prefecture := c.Query("prefecture")
    pageStr := c.Query("page")
    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        page = 1
    }
    limit := 10
    offset := (page - 1) * limit

    log.Printf("Received prefecture: %s, page: %d", prefecture, page)

    cities, err := fetchCities(prefecture, limit, offset)
    if err != nil {
        log.Printf("Failed to fetch cities: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{ "error": "Failed to fetch cities"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "Cities": cities,
        "Page":   page,
    })
}