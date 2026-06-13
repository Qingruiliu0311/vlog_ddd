package api

import (
	"net/http"

	"github.com/Qingruiliu0311/vlog_ddd/exception"
	"github.com/Qingruiliu0311/vlog_ddd/token"
	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	TokenService token.Service
}

func NewTokenHandler() *TokenHandler {
	return &TokenHandler{
		TokenService: token.GetService(),
	}
}

func (th *TokenHandler) IssueToken(ctx *gin.Context) {
	req := token.IssueTokenReq{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exception.NewBadRequest("Token request serialization failed %s", err))
		return
	}
	result, err := th.TokenService.IssueToken(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.NewInternalServerRequest("cannot issue token %s", err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (th *TokenHandler) ValidateToken(ctx *gin.Context) {
	req := token.ValidateTokenReq{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, exception.NewBadRequest("Json serialization failed %s", err))
		return
	}
	result, err := th.TokenService.ValidateToken(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, exception.NewInternalServerRequest("token validation failed %s", err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}
