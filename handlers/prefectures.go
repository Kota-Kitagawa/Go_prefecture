package handlers

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "Go_prefecture/internal/pkg"
)

func PrefectureHandler(c *gin.Context) {
    prefectures, err := pkg.FetchPrefecture()
    if err != nil {
        log.Printf("Failed to fetch prefectures: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch prefectures")
        return
    }
    responseFormat := c.Query("format")
    res :=pkg.GetResponse(responseFormat,"prefectures.html")
    res.Respond(c,gin.H{
        "Prefectures": prefectures,
    })
}


