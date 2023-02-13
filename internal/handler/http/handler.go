package http

import (
	"auction/internal/config"
	v1 "auction/internal/handler/http/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(corsMiddleware)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler()
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
