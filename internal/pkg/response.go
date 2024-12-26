package pkg
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// レスポンスの生成のためのインターフェース
type ResponseStrategy interface {
	Respond(c *gin.Context,data gin.H)
}

// JSON形式のレスポンスを返す構造体
type JSONResponse struct{}
func (j *JSONResponse) Respond(c *gin.Context,data gin.H){
	c.JSON(http.StatusOK,data)
}

// HTML形式のレスポンスを返す構造体
type HTMLResponse struct{
	TemplateName string
}
func (h *HTMLResponse) Respond(c *gin.Context,data gin.H){
	c.HTML(http.StatusOK,h.TemplateName,data)
}