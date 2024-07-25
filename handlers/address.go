package handlers

import (
	"net/http"
	"GO_prefecture/models"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html",gin.H{})
}
func SearchHandler(c *gin.Context) {
	prefecture := c.PostForm("prefecture")
	city := c.PostForm("city")
	town := c.PostForm("town")
	street := c.PostForm("street")
	postalCode,err := models.GetPostalCode(prefecture,city,town,street)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK.gin.H("postal_code" : postalCode))
}