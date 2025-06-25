package url

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type URL struct {
    ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
    LongURL   string    `gorm:"type:text;not null" json:"long_url"`
    ShortCode string    `gorm:"type:varchar(10);uniqueIndex;not null" json:"short_code"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (u *URL) BeforeCreate(tx *gorm.DB) (err error) {
    u.ID = uuid.New()
    return
}
