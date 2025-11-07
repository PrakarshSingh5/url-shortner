package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/PrakarshSingh5/url-shortner/backend/internal/repository"
)

// Redirect handles GET /:slug requests and redirects to the original URL.
func (h *Handler) Redirect(c *gin.Context) {
    slug := c.Param("slug")
    if slug == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Slug is required"})
        return
    }

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
