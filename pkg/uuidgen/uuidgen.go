package uuidgen

import (
    "strings"

    "github.com/google/uuid"
)

// GenerateShortCode возвращает короткий уникальный код из первых 8 символов UUID
func GenerateShortCode() string {
    id := uuid.New()
    return strings.ReplaceAll(id.String()[0:8], "-", "")
}
