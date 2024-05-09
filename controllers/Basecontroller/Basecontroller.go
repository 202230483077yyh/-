package Basecontroller

import "github.com/gin-gonic/gin"

type Basecontroller struct{}

func (con Basecontroller) Success(c *gin.Context) {
	c.String(200, "成功")
}

func (con Basecontroller) Fail(c *gin.Context) {
	c.String(200, "失败")
}
