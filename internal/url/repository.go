package url

import (
    "errors"

    "gorm.io/gorm"
)

type Repository interface {
    Create(url *URL) error
    FindByShortCode(shortCode string) (*URL, error)
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
    return &repository{db: db}
}

func (r *repository) Create(url *URL) error {
    return r.db.Create(url).Error
}

func (r *repository) FindByShortCode(shortCode string) (*URL, error) {
    var url URL
    result := r.db.Where("short_code = ?", shortCode).First(&url)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &url, result.Error
}
