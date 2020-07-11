package jwt

import (
	"blog/pkg/e"
	"blog/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			clamins, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else {
				if time.Now().Unix() > clamins.ExpiresAt {
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				}
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
