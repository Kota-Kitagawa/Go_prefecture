package handlers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"Go_prefecture/database"
)

func AddressHandler(c *gin.Context) {
	postalCode := c.Query("postalCode")
	if postalCode == "" {
		c.String(http.StatusBadRequest, "Postal code not specified")
		return
	}

	var address struct {
		Prefecture string
		City       string
		Address    string
	}

	err := database.DB.QueryRow("SELECT field7, field8, field9 FROM addresses WHERE field3 = ?", postalCode).Scan(
		&address.Prefecture, &address.City, &address.Address)
	if err != nil {
		log.Printf("Failed to fetch address: %v", err)
		c.String(http.StatusInternalServerError, "Failed to fetch address")
		return
	}

	c.HTML(http.StatusOK, "address.html", address)
}
