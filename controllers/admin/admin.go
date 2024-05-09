package admin

import (
	"Basecontroller"

	"github.com/gin-gonic/gin"
)

type Usercontroller struct {
	Basecontroller.Basecontroller
}

func (con Usercontroller) Index(c *gin.Context) {
	c.String(200, "这是首页")
}
