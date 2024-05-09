package routers

import "github.com/gin-gonic/gin"
import "admin"

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", admin.Usercontroller{}.Index)
		// adminRouters.GET("/user", admin.Usercontroller{}.Add)
		// adminRouters.GET("/article", admin.Usercontroller{}.News)
	}
}
