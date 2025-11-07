package utils

import (
    "crypto/rand"
    "math/big"
    "net/url"
    "strings"
)

const (
    slugAlphabet       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    defaultSlugLength  = 6
    secureDefaultProto = "https://"
)

// IsValidURL validates whether the provided string is an HTTP or HTTPS URL.
func IsValidURL(value string) bool {
    parsed, err := url.ParseRequestURI(value)
    if err != nil {
        return false
    }

    if parsed.Scheme != "http" && parsed.Scheme != "https" {
        return false
    }

    return parsed.Host != ""
}

// GenerateSlug returns a random slug of the provided length using a secure random source.
func GenerateSlug(length int) string {
    if length <= 0 {
        length = defaultSlugLength
    }

    builder := strings.Builder{}
    builder.Grow(length)

    alphabetLength := big.NewInt(int64(len(slugAlphabet)))
    for i := 0; i < length; i++ {
        index, err := rand.Int(rand.Reader, alphabetLength)
        if err != nil {
            // Fallback to first character to avoid failing completely
            builder.WriteByte(slugAlphabet[0])
            continue
        }
        builder.WriteByte(slugAlphabet[index.Int64()])
    }

    return builder.String()
}

// NormalizeURL ensures the URL has a protocol; defaults to https if absent.
func NormalizeURL(raw string) string {
    trimmed := strings.TrimSpace(raw)
    if trimmed == "" {
        return trimmed
    }

    if strings.HasPrefix(trimmed, "http://") || strings.HasPrefix(trimmed, "https://") {
        return trimmed
    }

    return secureDefaultProto + trimmed
}
