package utils

import "github.com/gin-gonic/gin"

func BindingReqBody(c *gin.Context, body any) error {
	return c.ShouldBindJSON(body)
}
