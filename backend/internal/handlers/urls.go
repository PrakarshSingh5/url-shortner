package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type urlResponse struct {
	ID          int64  `json:"id"`
	Slug        string `json:"slug"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	CreatedAt   string `json:"created_at"`
}

// ListURLs handles GET /api/urls requests.
func (h *Handler) ListURLs(c *gin.Context) {
	urls, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch URLs"})
		return
	}

	response := make([]urlResponse, 0, len(urls))
	for _, url := range urls {
		response = append(response, urlResponse{
			ID:          url.ID,
			Slug:        url.Slug,
			OriginalURL: url.OriginalURL,
			ShortURL:    h.cfg.BaseURL + "/" + url.Slug,
			CreatedAt:   url.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}
