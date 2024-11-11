package handlers

import (
	"github.com/gin-gonic/gin"
)

func AddressSearchHandler(c *gin.Context){
	c.HTML(200,"postsearch.html",nil)
}