package handlers

import (
	"net/http"
	"Go_prefecture/database"
	"github.com/gin-gonic/gin"
)

func CitiesHandler(c *gin.Context) {
	prefecture := c.PostForm("prefecture")

	rows, err := database.DB.Query("SELECT DISTINCT city FROM addresses WHERE prefecture = ?", prefecture)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var cities []string
	for rows.Next() {
		var city string
		if err := rows.Scan(&city); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		cities = append(cities, city)
	}

	c.HTML(http.StatusOK, "cities.html", gin.H{"prefecture": prefecture, "cities": cities})
}
