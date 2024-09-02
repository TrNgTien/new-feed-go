package utils

import "github.com/gin-gonic/gin"

func BindingBody(c *gin.Context, body any) error {
	return c.ShouldBindJSON(body)
}
