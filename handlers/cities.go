package handlers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"Go_prefecture/database"
)

func CityHandler(c *gin.Context) {
	prefecture := c.Query("prefecture")
	if prefecture == "" {
		c.String(http.StatusBadRequest, "Prefecture not specified")
		return
	}

	rows, err := database.DB.Query("SELECT DISTINCT field8 FROM addresses WHERE field7 = ?", prefecture)
	if err != nil {
		log.Printf("Failed to fetch cities: %v", err)
		c.String(http.StatusInternalServerError, "Failed to fetch cities")
		return
	}
	defer rows.Close()

	var cities []string
	for rows.Next() {
		var city string
		rows.Scan(&city)
		cities = append(cities, city)
	}

	c.HTML(http.StatusOK, "citiesresult.html", gin.H{
		"Prefecture": prefecture,
		"Cities":     cities,
	})
}