package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//untuk output json
//biar ambil format warper gak usang ngetik ulang2
func WarpAPIError(c * gin.Context, message string, code int){
c.JSON(code,map[string]interface{}{
	"code": code,
	"error_type" : http.StatusText(code),
	"error_detail": message,
})
}

func WarpAPISucces(c * gin.Context, message string, code int){
	c.JSON(code,map[string]interface{}{
		"code":   code,
		"status": message,
	})
}


func WarpAPIData(c * gin.Context, data interface{}, code int, message string){
	c.JSON(code,map[string]interface{}{
		"code":   code,
		"status": message,
		"data": data,
	})
}
