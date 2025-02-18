package stocks

import (
	"net/http"

	v1 "github.com/Gaurav-coding08/ingestion-go/pkg/client"
	"github.com/gin-gonic/gin"
)

type service interface {
	UpdateStockPrice(
		updateStockPrice v1.UpdateStockPrice,
	) error
}

type Controller struct {
	service service
}

func New(service service) *Controller {
	return &Controller{
		service: service}
}

func (ctrl *Controller) Update(c *gin.Context) {
	updateReq := v1.UpdateStockPrice{}

	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.UpdateStockPrice(updateReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated,"")

}
