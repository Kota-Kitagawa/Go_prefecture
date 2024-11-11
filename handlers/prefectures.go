package handlers

import (
	"log"
	"net/http"
	"regexp"
	"github.com/gin-gonic/gin"
	"Go_prefecture/database"
)
func PrefectureHandler(c *gin.Context) {
	rows, err := database.DB.Query("SELECT DISTINCT field7 FROM addresses")
	if err != nil {
		log.Printf("Failed to fetch prefectures: %v", err)
		c.String(http.StatusInternalServerError, "Failed to fetch prefectures")
		return
	}
	defer rows.Close()
	// 正規表現のコードを二回使っている上のSELEｃT分と下のMuxtCompileの文

	var prefectures []string
	reader:=regexp.MustCompile(`^[\p{Han}]{2,3}(?:都|道|府|県)$`)

	for rows.Next() {
		var prefecture string
		rows.Scan(&prefecture)
		if(reader.MatchString(prefecture)){
			prefectures = append(prefectures, prefecture)
		}
	}
	c.HTML(http.StatusOK, "prefectures.html", gin.H{
		"Prefectures": prefectures,
	})
}
