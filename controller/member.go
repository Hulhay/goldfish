package controller

import (
	"net/http"

	"github.com/Hulhay/goldfish/shared"
	"github.com/Hulhay/goldfish/usecase"
	"github.com/Hulhay/goldfish/usecase/member"
	"github.com/gin-gonic/gin"
)

type memberController struct {
	memberUC usecase.Member
}

type MemberContoller interface {
	InsertMember(ctx *gin.Context)
	GetMember(ctx *gin.Context)
}

func NewMemberController(memberUC usecase.Member) MemberContoller {
	return &memberController{
		memberUC: memberUC,
	}
}

func (c *memberController) InsertMember(ctx *gin.Context) {

	var (
		params member.InsertMemberRequest
		err    error
	)

	err = ctx.ShouldBind(&params)
	if err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = c.memberUC.InsertMember(ctx, params)
	if err != nil {
		res := shared.BuildErrorResponse("Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := shared.BuildResponse("Success!", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *memberController) GetMember(ctx *gin.Context) {

	var (
		params   member.GetMemberRequest
		response []member.MemberListResponse
		err      error
	)

	params.MemberNIK = ctx.Query("member_nik")

	response, err = c.memberUC.GetMember(ctx, params)
	if err != nil {
		res := shared.BuildErrorResponse("Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := shared.BuildResponse("Success!", response)
	ctx.JSON(http.StatusOK, res)
}
