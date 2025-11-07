package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"

    "github.com/PrakarshSingh5/url-shortner/backend/internal/repository"
    "github.com/PrakarshSingh5/url-shortner/backend/internal/utils"
)

const maxSlugAttempts = 10

type shortenRequest struct {
    URL string `json:"url"`
}

// ShortenURL handles POST /api/shorten requests.
func (h *Handler) ShortenURL(c *gin.Context) {
    var req shortenRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if req.URL == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
        return
    }

    normalized := utils.NormalizeURL(req.URL)
    if !utils.IsValidURL(normalized) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
        return
    }

    slug := utils.GenerateSlug(0)
    for attempts := 0; attempts < maxSlugAttempts; attempts++ {
        exists, err := h.repo.SlugExists(slug)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check slug availability"})
            return
        }

        if !exists {
            break
        }
        slug = utils.GenerateSlug(0)

        if attempts == maxSlugAttempts-1 {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate unique slug"})
            return
        }
    }

    url, err := h.repo.Create(slug, normalized)
    if err != nil {
        if err == repository.ErrDuplicateSlug {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate unique slug"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL"})
        return
    }

    shortURL := h.cfg.BaseURL + "/" + url.Slug

    c.JSON(http.StatusOK, gin.H{
        "id":           url.ID,
        "slug":         url.Slug,
        "original_url": url.OriginalURL,
        "short_url":    shortURL,
        "created_at":   url.CreatedAt.Format(time.RFC3339),
    })
}
