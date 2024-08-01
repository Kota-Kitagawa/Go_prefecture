package handlers

import (
	"github.com/gin-gonic/gin"
)

func PostalHandler(c *gin.Context) {
	c.HTML(200, "postcode.html", nil)
}
