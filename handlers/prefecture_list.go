package handlers

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "Go_prefecture/internal/pkg"
)

func PrefListHTMLHandler(c *gin.Context) {
    prefectures, err := pkg.FetchPrefecture()
    if err != nil {
        log.Printf("Failed to fetch prefectures: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch prefectures")
        return
    }
	responseFormat := "html"
    res :=pkg.GetResponse(responseFormat,"cities.html")
    res.Respond(c,gin.H{
        "Prefectures": prefectures,
    })
}