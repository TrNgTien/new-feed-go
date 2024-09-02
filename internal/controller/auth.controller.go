package controller

import (
	"fmt"

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

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var lp LoginPayload
	if err := utils.BindingBody(c, &lp); err != nil {
		fmt.Println("[Login] BindingBody error: ", &lp)
		response.R400(c)
		return
	}

	user, err := ac.Service.Login(lp.Username, lp.Password)
	if err != nil {
		response.R400(c)
		return
	}

	response.R200(c, user)
}

func (ac *AuthController) Register(c *gin.Context) {
	var lp LoginPayload
	if err := utils.BindingBody(c, &lp); err != nil {
		response.R400(c)
		return
	}

	isCreated, err := ac.Service.Register(lp.Username, lp.Password)

	if err != nil {
		response.R400(c)
		return
	}

	response.R200(c, isCreated)
}

func (ac *AuthController) Logout(c *gin.Context) {
	response.R200(c, true)
}
