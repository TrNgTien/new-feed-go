package controller

import (
	"github.com/TrNgTien/new-feed-go/internal/helpers/response"
	"github.com/TrNgTien/new-feed-go/internal/services"
	"github.com/TrNgTien/new-feed-go/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{Service: service}
}

type userPayload struct {
	Username string `json:"username" validate:"required" errorMsg:"Username is required"`
	Password string `json:"password" validate:"required" errorMsg:"Password is required"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var up userPayload

	if err := utils.BindingReqBody(c, &up); err != nil {
		response.R400(c, err.Error())
		return
	}

	if err := utils.Validate.Struct(&up); err != nil {
		response.R400(c, utils.ParseErrorMsg(up, err))
		return
	}

	user, err := ac.Service.Login(up.Username, up.Password)
	if err != nil {
		response.R400(c, err.Error())
		return
	}

	response.R200(c, user)
}

func (ac *AuthController) Register(c *gin.Context) {
	var up userPayload
	if err := utils.BindingReqBody(c, &up); err != nil {
		response.R400(c, err.Error())
		return
	}

	if err := utils.Validate.Struct(&up); err != nil {
		response.R400(c, utils.ParseErrorMsg(up, err))
		return
	}

	isCreated, err := ac.Service.Register(up.Username, up.Password)

	if err != nil {
		response.R400(c, err.Error())
		return
	}

	response.R200(c, isCreated)
}

func (ac *AuthController) Logout(c *gin.Context) {
	response.R200(c, true)
}
