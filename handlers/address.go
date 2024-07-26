package handlers

import (
	"net/http"
	"Go_prefecture/database"
	"Go_prefecture/models"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	rows, err := database.DB.Query("SELECT DISTINCT prefecture FROM addresses")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var prefectures []string
	for rows.Next() {
		var prefecture string
		if err := rows.Scan(&prefecture); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		prefectures = append(prefectures, prefecture)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{"prefectures": prefectures})
}

func SearchHandler(c *gin.Context) {
	prefecture := c.PostForm("prefecture")
	city := c.PostForm("city")
	town := c.PostForm("town")
	street := c.PostForm("street")

	postalCode, err := models.GetPostalCode(prefecture, city, town, street)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"postal_code": postalCode})
}
