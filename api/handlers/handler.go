package handlers

import (
	"strconv"

	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/clients"
	"github.com/AbdulahadAbduqahhorov/gRPC/blogpost/api_gateway/config"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *clients.GrpcClients
	cfg config.Config
}

func NewHandler(s *clients.GrpcClients,cfg config.Config) Handler {
	return Handler{
		services:s,
		cfg: cfg,
	}
}

func (h *Handler) getOffsetParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getLimitParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}
