package lunch

import (
	"bc-labs-lunch-bot/di"
	"bc-labs-lunch-bot/internal/lunch/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	lunchService service.LunchService
)

func init() {
	lunchService = di.InitializeLunchServiceImpl()
}

func CreateModalLunchCosts(c *gin.Context) {
	err := lunchService.RegisterLunchPayment(*c.Request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}
