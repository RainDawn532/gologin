package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseSuccessful(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"data": msg,
	})
}

func ResponseWrong(c *gin.Context, errorCode int, errorMsg string) {
	c.JSON(errorCode, map[string]interface{}{
		"code":   errorCode,
		"errMsg": errorMsg,
	})
}
