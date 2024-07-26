package handlers

import (
	"net/http"
	"Go_prefecture/database"
	"github.com/gin-gonic/gin"
)

func TownsHandler(c *gin.Context) {
	prefecture := c.PostForm("prefecture")
	city := c.PostForm("city")

	rows, err := database.DB.Query("SELECT DISTINCT town FROM addresses WHERE prefecture = ? AND city = ?", prefecture, city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var towns []string
	for rows.Next() {
		var town string
		if err := rows.Scan(&town); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		towns = append(towns, town)
	}

	c.HTML(http.StatusOK, "towns.html", gin.H{"prefecture": prefecture, "city": city, "towns": towns})
}
