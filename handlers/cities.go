package handlers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"Go_prefecture/database"
)

func CityHandler(c *gin.Context) {
	prefecture := c.PostForm("prefecture")
	log.Printf("Received prefecture: %s", prefecture)

	query := `
		SELECT CASE
			WHEN field9 = '以下に掲載がない場合' THEN field8
			ELSE field8 || field9
			END AS city
		FROM addresses
		WHERE field7 = ?
		`
	
	rows, err := database.DB.Query(query, prefecture)
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
		"Cities":  cities,
	})
}