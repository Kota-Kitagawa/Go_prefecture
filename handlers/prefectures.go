package handlers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"Go_prefecture/pkg/database"
)
func PrefectureHandler(c *gin.Context) {
	rows, err := database.DB.Query("SELECT DISTINCT field7 FROM addresses")
	if err != nil {
		log.Printf("Failed to fetch prefectures: %v", err)
		c.String(http.StatusInternalServerError, "Failed to fetch prefectures")
		return
	}
	defer rows.Close()
	
	var prefectures []string
	for rows.Next() {
		var prefecture string
		if err:=rows.Scan(&prefecture); err != nil {
			log.Printf("Failed to scan prefecture: %v", err)
			c.String(http.StatusInternalServerError, "Failed to scan prefecture")
			return
		}
		prefectures = append(prefectures, prefecture)
	}
	c.HTML(http.StatusOK, "prefectures.html", gin.H{
		"Prefectures": prefectures,
	})
}
