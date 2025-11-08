package repository

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/PrakarshSingh5/url-shortner/backend/internal/models"
)

const sqliteDateLayout = "2006-01-02 15:04:05"

var (
	// ErrNotFound indicates the requested slug does not exist in the database.
	ErrNotFound = errors.New("url not found")
	// ErrDuplicateSlug indicates the slug already exists in the database.
	ErrDuplicateSlug = errors.New("slug already exists")
)

// URLRepository provides persistence helpers for URL entities.
type URLRepository struct {
	db *sql.DB
}

// NewURLRepository constructs a new URLRepository instance.
func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

// Create inserts a new shortened URL into the database.
func (r *URLRepository) Create(slug, originalURL string) (*models.URL, error) {
	query := `INSERT INTO urls (slug, original_url) VALUES (?, ?)`

	result, err := r.db.Exec(query, slug, originalURL)
	if err != nil {
		if isUniqueConstraintError(err) {
			return nil, ErrDuplicateSlug
		}
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	url, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	return url, nil
}

// Fetch by original URL
func (r *URLRepository) GetByOriginalURL(originalUrl string) (*models.URL, error) {
	const query = `SELECT id, slug,original_url, created_at FROM urls WHERE original_url=? LIMIT 1`
	row := r.db.QueryRow(query, originalUrl)

	var url models.URL
	if err := row.Scan(&url.ID, &url.Slug, &url.OriginalURL, &url.CreatedAt); err != nil {
		if err == sql.ErrNoRows {

			return nil, nil
		}
		log.Printf("ERROR in GetByOriginalURL query: %v", err)
		return nil, err
	}
	return &url, nil
}

// GetByID fetches a URL by its numeric identifier.
func (r *URLRepository) GetByID(id int64) (*models.URL, error) {
	query := `SELECT id, slug, original_url, created_at FROM urls WHERE id = ?`

	row := r.db.QueryRow(query, id)
	return scanURL(row)
}

// GetBySlug fetches a URL by its slug.
func (r *URLRepository) GetBySlug(slug string) (*models.URL, error) {
	query := `SELECT id, slug, original_url, created_at FROM urls WHERE slug = ?`

	row := r.db.QueryRow(query, slug)
	return scanURL(row)
}

// GetAll returns all shortened URLs ordered by creation date descending.
func (r *URLRepository) GetAll() ([]models.URL, error) {
	query := `SELECT id, slug, original_url, created_at FROM urls ORDER BY datetime(created_at) DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	urls := make([]models.URL, 0)
	for rows.Next() {
		var model models.URL
		var createdAt string

		if err := rows.Scan(&model.ID, &model.Slug, &model.OriginalURL, &createdAt); err != nil {
			return nil, err
		}

		model.CreatedAt = createdAt
		urls = append(urls, model)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

// SlugExists checks whether a slug is already present in the database.
func (r *URLRepository) SlugExists(slug string) (bool, error) {
	query := `SELECT COUNT(1) FROM urls WHERE slug = ?`

	var count int
	if err := r.db.QueryRow(query, slug).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func scanURL(row *sql.Row) (*models.URL, error) {
	var model models.URL
	var createdAt string

	if err := row.Scan(&model.ID, &model.Slug, &model.OriginalURL, &createdAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	
	model.CreatedAt = createdAt

	return &model, nil
}

func isUniqueConstraintError(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "unique constraint failed")
}
