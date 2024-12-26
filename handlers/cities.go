package handlers

import (
    "log"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "Go_prefecture/internal/pkg"
)

func CitiesHandler(c *gin.Context) {
    prefecture := c.Query("prefecture")
    pageStr := c.Query("page")
    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        page = 1
    }
    limit := 10
    offset := (page - 1) * limit
    log.Printf("Received prefecture: %s, page: %d", prefecture, page)
    cities, err := pkg.FetchCities(prefecture, limit, offset)
    if err != nil {
        log.Printf("Failed to fetch cities: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch cities")
        return
    }
    responseFormat := c.Query("format")
    res := pkg.GetResponse(responseFormat,"citiesresult.html")
    res.Respond(c,gin.H{
        "Cities": cities,
        "Page":   page,
        "Prefecture": prefecture,
    })
}

