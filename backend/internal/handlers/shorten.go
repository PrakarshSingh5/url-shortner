package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/PrakarshSingh5/url-shortner/backend/internal/database"
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

	cachedSlug, err := database.Rdb.Get(database.Ctx, normalized).Result()
	if err == nil {
		// Found in Redis — return it immediately
		shortURL := h.cfg.BaseURL + "/" + cachedSlug
		c.JSON(http.StatusOK, gin.H{
			"slug":         cachedSlug,
			"original_url": normalized,
			"short_url":    shortURL,
			"cached":       true,
			"message":      "Fetched from cache",
		})
		return
	}

	// 2 Step 2: If not found in Redis, check DB if URL already exists
	existing, err := h.repo.GetByOriginalURL(normalized)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing URLs"})
		return
	}

	if existing != nil {
		// Found in DB → store it in Redis for next time & return it
		_ = database.Rdb.Set(database.Ctx, normalized, existing.Slug, 24*time.Hour).Err()
		shortURL := h.cfg.BaseURL + "/" + existing.Slug
		c.JSON(http.StatusOK, gin.H{
			"id":           existing.ID,
			"slug":         existing.Slug,
			"original_url": existing.OriginalURL,
			"short_url":    shortURL,
			"created_at":   existing.CreatedAt,
			"cached":       false,
			"message":      "Fetched from DB",
		})
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

	// Add this new Redis caching part here

	_ = database.Rdb.Set(database.Ctx, slug, normalized, 24*time.Hour).Err() // slug → url (for redirects)
	_ = database.Rdb.Set(database.Ctx, normalized, slug, 24*time.Hour).Err()

	shortURL := h.cfg.BaseURL + "/" + url.Slug

	c.JSON(http.StatusOK, gin.H{
		"id":           url.ID,
		"slug":         url.Slug,
		"original_url": url.OriginalURL,
		"short_url":    shortURL,
		"created_at":   url.CreatedAt,
	})
}
