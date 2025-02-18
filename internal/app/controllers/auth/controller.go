package auth

import (
	"net/http"

	"github.com/Gaurav-coding08/ingestion-go/internal/app/models"
	v1 "github.com/Gaurav-coding08/ingestion-go/pkg/client"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Create(
		userRegisterReq v1.CreateUserRequest,
	) (models.User, error)
	Login(loginUser models.LoginUser) (*models.AuthToken, error)
}

type Controller struct {
	authService AuthService
}

func New(service AuthService) *Controller {
	return &Controller{
		authService: service}
}

func (ctrl *Controller) Register(c *gin.Context) {
	var userRegisterreq v1.CreateUserRequest

	if err := c.ShouldBindJSON(&userRegisterreq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.authService.Create(userRegisterreq)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user.ToResponse())
}

func (ctrl *Controller) Login(c *gin.Context) {
	var userLoginreq = v1.LoginUserRequest{}

	if err := c.ShouldBindJSON(&userLoginreq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	authToken, err := ctrl.authService.Login(models.LoginUser{}.FromRequest(userLoginreq))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})

		return
	}

	c.JSON(http.StatusOK, authToken.ToResponse())
}
