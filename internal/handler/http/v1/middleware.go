package v1

import (
	"auction/internal/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "uuid"
	roleCtx             = "role"
)

func (h *Handler) parseAuthHeader(c *gin.Context) (model.TokenBody, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return model.TokenBody{}, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return model.TokenBody{}, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return model.TokenBody{}, errors.New("token is empty")
	}

	return h.tokenManager.ParseJWT(headerParts[1])
}

func (h *Handler) userIdentity(c *gin.Context) {
	body, err := h.parseAuthHeader(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, body.Uuid)
	c.Set(roleCtx, body.Role)
}

func (h *Handler) staffIdentity(c *gin.Context) {
	body, err := h.parseAuthHeader(c)
	if err != nil || body.Role == "user" {
		newResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, body.Uuid)
}

func (h *Handler) ownerIdentity(c *gin.Context) {
	body, err := h.parseAuthHeader(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
	} else if body.Role != "owner" {
		newResponse(c, http.StatusUnauthorized, "permission denied")
	}

	c.Set(userCtx, body.Uuid)
}
