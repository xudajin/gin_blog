package routers

import (
	"blog/middleware/jwt"
	"blog/pkg/setting"
	"blog/routers/api"
	v1 "blog/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.RunMode)
	r.POST("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1", jwt.Jwt())
	{ // 标签路由
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tag/:id", v1.EditTag)
		apiv1.DELETE("/tag/:id", v1.DeleteTag)
		// 文章路由
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/article/:id", v1.GetArticles)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/article/:id", v1.EditArticle)
		apiv1.DELETE("/article/:id", v1.DeleteArticle)

	}

	return r
}
