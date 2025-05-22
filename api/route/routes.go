package route

import (
	"github.com/arvinpaundra/go-boilerplate/application/resthttp"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(g *gin.Engine, db *gorm.DB) {
	controller := resthttp.NewController(db)

	g.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
	}))

	g.Use(gin.Recovery())

	v1 := g.Group("/api/v1")

	helloWorld := v1.Group("/hello-world")
	helloWorld.GET("", controller.GetHelloWorld)
}
