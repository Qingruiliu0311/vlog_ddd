package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Qingruiliu0311/vlog_ddd/token"
	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context) {
	accessToken := ctx.GetHeader("Authorization")
	akPair := strings.Split(accessToken, " ")
	if len(akPair) > 1 {
		accessToken = akPair[1]
	}
	//验证token
	req := token.ValidateTokenReq{
		AccessToken: accessToken,
	}
	token, err := token.GetService().ValidateToken(ctx.Request.Context(), &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c := context.WithValue(ctx.Request.Context(), TokenCtxKey{}, token)
	ctx.Request = ctx.Request.WithContext(c)
}

type TokenCtxKey struct{}
