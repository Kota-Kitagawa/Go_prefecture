package handlers

import (
	"github.com/gin-gonic/gin"
	"Go_prefecture/pkg"
)

func HomeHandler(c *gin.Context) {
	responseFormat := "html"
	res := pkg.GetResponse(responseFormat,"index.html")
	res.Respond(c,gin.H{})
}
