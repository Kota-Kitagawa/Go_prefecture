package handlers

import (
	"github.com/gin-gonic/gin"
	"Go_prefecture/pkg"
)

func AddressSearchHandler(c *gin.Context) {
	responseFormat := "html"
	res := pkg.GetResponse(responseFormat,"postsearch.html")
	res.Respond(c,gin.H{})
}