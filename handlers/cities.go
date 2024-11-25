package handlers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"Go_prefecture/pkg/database"
)

func fetchCities(prefecture string)([]string,error){
	query := `
		SELECT CASE
			WHEN field9 = '以下に掲載がない場合' THEN field8
			ELSE field8 || field9
			END AS city
		FROM addresses
		WHERE field7 = ?
		`
	
	rows, err := database.DB.Query(query, prefecture)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
}
func CitiesHTMLHandler(c *gin.Context){
	prefecture := c.PostForm("prefecture")
	log.Printf("Received prefecture: %s", prefecture)

	cities, err := fetchCities(prefecture)
	if err != nil{
		log.Printf("Failed to fetch cities: %v", err)
		c.String(http.StatusInternalServerError, "Failed to fetch cities")
		return
	}

	c.HTML(http.StatusOK, "citiesresult.html", gin.H{
		"Cities": cities,
	})
}

func CitiesJSONHandler(c *gin.Context){
	prefecture := c.PostForm("prefecture")
	log.Printf("Received prefecture: %s", prefecture)

	cities, err := fetchCities(prefecture)
	if err != nil{
		log.Printf("Failed to fetch cities: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{ "error": "Failed to fetch cities"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Cities": cities,
	})
}