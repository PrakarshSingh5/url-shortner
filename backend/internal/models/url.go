package models

// URL represents a shortened URL record stored in the database.
type URL struct {
	ID          int64  `json:"id"`
	Slug        string `json:"slug"`
	OriginalURL string `json:"original_url"`
	CreatedAt   string `json:"created_at"`
}
