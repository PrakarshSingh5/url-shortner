package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PrakarshSingh5/url-shortner/backend/internal/database"
	"github.com/PrakarshSingh5/url-shortner/backend/internal/repository"
)

// Redirect handles GET /:slug requests and redirects to the original URL.
func (h *Handler) Redirect(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Slug is required"})
		return
	}

	//first try in redis then go to check for the db
	cachedURL, err := database.Rdb.Get(database.Ctx, slug).Result()
	if err == nil {
		// Found in cache  redirect immediately

		c.Redirect(http.StatusFound, cachedURL)
		return
	}

    // if not found in redis then try for this
	url, err := h.repo.GetBySlug(slug)
	if err != nil {
		if err == repository.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to resolve slug"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, url.OriginalURL)
}
