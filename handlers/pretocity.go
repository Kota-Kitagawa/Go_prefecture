package handlers

import (
	"github.com/gin-gonic/gin"
)

func PrefecturetocityHandler(c *gin.Context) {
	c.HTML(200, "cities.html", nil)
}