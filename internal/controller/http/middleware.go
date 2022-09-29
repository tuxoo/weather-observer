package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tuxoo/idler/pkg/auth"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	id, err := h.parseAuthHeader(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
	}

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.Set(userCtx, id)
}

func (h *Handler) parseAuthHeader(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.ParseToken(auth.Token(headerParts[1]))
}

func getUserId(c *gin.Context) (id string, err error) {
	ctxId, ok := c.Get(userCtx)
	if !ok {
		err = errors.New("user doesn't exist in context")
	}

	id = ctxId.(string)
	return
}
