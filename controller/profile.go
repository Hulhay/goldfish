package controller

import (
	"net/http"

	"github.com/Hulhay/goldfish/shared"
	"github.com/Hulhay/goldfish/usecase"
	"github.com/gin-gonic/gin"
)

type profileController struct {
	profileUC usecase.Profile
}

type ProfileController interface {
	GetProfile(ctx *gin.Context)
}

func NewProfileController(profileUC usecase.Profile) ProfileController {
	return &profileController{
		profileUC: profileUC,
	}
}

func (c *profileController) GetProfile(ctx *gin.Context) {

	email := ctx.GetString("email")

	response, err := c.profileUC.GetProfile(ctx, email)
	if err != nil {
		res := shared.BuildErrorResponse("Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := shared.BuildResponse("Success!", response)
	ctx.JSON(http.StatusOK, res)
}
