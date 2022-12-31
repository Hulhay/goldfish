package controller

import (
	"net/http"

	"github.com/Hulhay/goldfish/model"
	"github.com/Hulhay/goldfish/shared"
	"github.com/Hulhay/goldfish/usecase"
	"github.com/Hulhay/goldfish/usecase/category"
	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryUC usecase.Category
}

type CategoryController interface {
	InsertCategory(ctx *gin.Context)
	GetListCategory(ctx *gin.Context)
}

func NewCategoryController(categoryUC usecase.Category) CategoryController {
	return &categoryController{
		categoryUC: categoryUC,
	}
}

func (c *categoryController) InsertCategory(ctx *gin.Context) {

	var (
		params category.InsertCategoryRequest
		err    error
	)

	err = ctx.ShouldBind(&params)
	if err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = c.categoryUC.InsertCategory(ctx, params)
	if err != nil {
		res := shared.BuildErrorResponse("Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := shared.BuildResponse("Success!", nil)
	ctx.JSON(http.StatusOK, res)

}

func (c *categoryController) GetListCategory(ctx *gin.Context) {

	var (
		response []model.Category
		err      error
	)

	response, err = c.categoryUC.GetListCategory(ctx)
	if err != nil {
		res := shared.BuildErrorResponse("Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := shared.BuildResponse("Success!", response)
	ctx.JSON(http.StatusOK, res)

}
