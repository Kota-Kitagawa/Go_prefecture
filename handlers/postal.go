package handlers

import (
	"github.com/gin-gonic/gin"
	"Go_prefecture/pkg"
)

func PostalHandler(c *gin.Context) {
	responseFormat := "html"
	res := pkg.GetResponse(responseFormat,"postcode.html")
	res.Respond(c,gin.H{})
}