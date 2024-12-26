package handlers

import (
    "log"
    "net/http"
    "Go_prefecture/pkg"
    "github.com/gin-gonic/gin"
)

func PretoCityHandler(c *gin.Context) {
    prefectures, err := pkg.FetchPretoCity()
    if err != nil {
        log.Printf("Failed to fetch prefectures: %v", err)
        c.String(http.StatusInternalServerError, "Failed to fetch prefectures")
        return
    }
    responseFormat := c.DefaultQuery("format","html")
    res :=pkg.GetResponse(responseFormat,"prefectures.html")
    res.Respond(c,gin.H{
        "Prefectures": prefectures,
    })
}

