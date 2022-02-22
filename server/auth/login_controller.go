package auth

import (
	"log"

	"github.com/andey-robins/alexa-skills-analytics/server/types"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService LoginService
	jwtService   JWTService
}

func LoginHandler(loginService LoginService, jwtService JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var creds types.LoginCredentials
	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		log.Println("Unable to bind creds:", err)
		return "no data found"
	}

	log.Println("Login creds:", creds)

	if isUserAuthenticated := controller.loginService.LoginWebUser(creds.Email, creds.Password); isUserAuthenticated {
		token := controller.jwtService.GenerateToken(creds.Email, true)
		return token
	}

	return ""
}
