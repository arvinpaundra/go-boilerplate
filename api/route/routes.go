package route

import (
	"github.com/arvinpaundra/go-boilerplate/api/route/auth"
	"github.com/arvinpaundra/go-boilerplate/application/resthttp"
	"github.com/arvinpaundra/go-boilerplate/core/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(g *gin.Engine, db *gorm.DB, vld *validator.Validator) {
	controller := resthttp.NewController(db, vld)

	g.Use(gin.Recovery())
	g.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
	}))

	PublicRoute(g, controller)
}

func PublicRoute(g *gin.Engine, cont *resthttp.Controller) {
	v1 := g.Group("/api/v1")

	auth.PublicRoute(v1, cont)
}
