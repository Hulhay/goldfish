package controller

import (
	"net/http"

	"github.com/Hulhay/goldfish/shared"
	"github.com/gin-gonic/gin"
)

type healthController struct{}

type HealthController interface {
	Health(ctx *gin.Context)
}

func NewHealthController() HealthController {
	return &healthController{}
}

func (c *healthController) Health(ctx *gin.Context) {
	resp := shared.BuildResponse("I'm feeling fine", nil)
	ctx.JSON(http.StatusOK, resp)
}
