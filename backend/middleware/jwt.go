package middleware

import (
	"github.com/gin-gonic/gin"
	"timedeliver/model/response"
	"timedeliver/util"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := util.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == util.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}
