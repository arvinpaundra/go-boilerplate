package resthttp

import (
	"net/http"

	"github.com/arvinpaundra/go-boilerplate/core/format"
	"github.com/arvinpaundra/go-boilerplate/core/token"
	"github.com/arvinpaundra/go-boilerplate/domain/auth/constant"
	"github.com/arvinpaundra/go-boilerplate/domain/auth/dto/request"
	"github.com/arvinpaundra/go-boilerplate/domain/auth/service"
	"github.com/arvinpaundra/go-boilerplate/infrastructure/auth"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (cont *Controller) Register(c *gin.Context) {
	var payload request.Register

	_ = c.ShouldBindJSON(&payload)

	verrs := cont.validator.Validate(payload)
	if verrs != nil {
		c.JSON(http.StatusBadRequest, format.BadRequest("invalid request body", verrs))
		return
	}

	handler := service.NewRegisterHandler(
		auth.NewUserReaderRepository(cont.db),
		auth.NewUserWriterRepository(cont.db),
	)

	err := handler.Handle(c, payload)
	if err != nil {
		switch err {
		case constant.ErrEmailAlreadyTaken:
			c.JSON(http.StatusConflict, format.Conflict(err.Error()))
			return
		default:
			c.JSON(http.StatusInternalServerError, format.InternalServerError())
			return
		}
	}

	c.JSON(http.StatusCreated, format.SuccessCreated("success register", nil))
}

func (cont *Controller) Login(c *gin.Context) {
	var payload request.Login

	_ = c.ShouldBindJSON(&payload)

	verrs := cont.validator.Validate(payload)
	if verrs != nil {
		c.JSON(http.StatusBadRequest, format.BadRequest("invalid request body", verrs))
		return
	}

	handler := service.NewLoginHandler(
		auth.NewUserReaderRepository(cont.db),
		auth.NewSessionWriterRepository(cont.db),
		auth.NewUserCacheRepository(cont.rdb),
		token.NewJWT(viper.GetString("JWT_SECRET")),
	)

	res, err := handler.Handle(c, payload)
	if err != nil {
		switch err {
		case constant.ErrUserNotFound, constant.ErrWrongEmailOrPassword:
			c.JSON(http.StatusUnprocessableEntity, format.UnprocessableEntity(constant.ErrWrongEmailOrPassword.Error()))
			return
		default:
			c.JSON(http.StatusInternalServerError, format.InternalServerError())
			return
		}
	}

	c.JSON(http.StatusOK, format.SuccessOK("success login", res))
}

func (cont *Controller) RefreshToken(c *gin.Context) {
	var payload request.RefreshToken

	_ = c.ShouldBindJSON(&payload)

	verrs := cont.validator.Validate(payload)
	if verrs != nil {
		c.JSON(http.StatusBadRequest, format.BadRequest("invalid request body", verrs))
		return
	}

	handler := service.NewRefreshTokenHandler(
		auth.NewUserReaderRepository(cont.db),
		auth.NewSessionReaderRepository(cont.db),
		auth.NewSessionWriterRepository(cont.db),
		auth.NewUserCacheRepository(cont.rdb),
		token.NewJWT(viper.GetString("JWT_SECRET")),
	)

	res, err := handler.Handle(c, payload)
	if err != nil {
		switch err {
		case constant.ErrUserNotFound, constant.ErrSessionNotFound, constant.ErrTokenInvalid:
			c.JSON(http.StatusUnprocessableEntity, format.UnprocessableEntity(err.Error()))
			return
		default:
			c.JSON(http.StatusInternalServerError, format.InternalServerError())
			return
		}
	}

	c.JSON(http.StatusCreated, format.SuccessCreated("success refresh token", res))
}
