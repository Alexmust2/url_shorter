package url

import (
    "errors"
    "url_shortener/pkg/uuidgen"
)

type Service interface {
    CreateShortURL(longURL string) (*URL, error)
    GetByShortCode(shortCode string) (*URL, error)
}

type service struct {
    repo Repository
}

func NewService(r Repository) Service {
    return &service{repo: r}
}

func (s *service) CreateShortURL(longURL string) (*URL, error) {
    if longURL == "" {
        return nil, errors.New("empty URL")
    }

    // Генерация короткого кода
    shortCode := uuidgen.GenerateShortCode()

    url := &URL{
        LongURL:   longURL,
        ShortCode: shortCode,
    }

    err := s.repo.Create(url)
    if err != nil {
        return nil, err
    }

    return url, nil
}

func (s *service) GetByShortCode(shortCode string) (*URL, error) {
    return s.repo.FindByShortCode(shortCode)
}
