package handlers

import (
    "github.com/gin-gonic/gin"

    "github.com/PrakarshSingh5/url-shortner/backend/internal/config"
    "github.com/PrakarshSingh5/url-shortner/backend/internal/repository"
)

// Handler aggregates dependencies for HTTP handlers.
type Handler struct {
    repo *repository.URLRepository
    cfg  *config.Config
}

// New constructs a Handler instance.
func New(repo *repository.URLRepository, cfg *config.Config) *Handler {
    return &Handler{repo: repo, cfg: cfg}
}

// RegisterRoutes wires all HTTP handlers.
func (h *Handler) RegisterRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.POST("/shorten", h.ShortenURL)
        api.GET("/urls", h.ListURLs)
    }

    router.GET("/:slug", h.Redirect)
}
