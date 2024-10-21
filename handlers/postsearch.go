package handlers

import (
	"log"
	"net/http"
	"regexp"
	"Go_prefecture/database"
	"github.com/gin-gonic/gin"
)

func PostSearchHandler(c *gin.Context){
	Prefecture :=c.PostForm("prefecture")

}