package resthttp

import (
	"net/http"

	"github.com/arvinpaundra/go-boilerplate/domain/helloworld/service"
	"github.com/gin-gonic/gin"
)

func (cont *Controller) GetHelloWorld(c *gin.Context) {
	command := service.ExampleCommand{}

	handler := service.NewExampleHandler()

	res, err := handler.Handle(c, command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"print": res,
	})
}
