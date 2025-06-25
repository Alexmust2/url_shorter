package url

import (
    "github.com/gofiber/fiber/v2"
)

type Handler struct {
    service Service
}

func NewHandler(s Service) *Handler {
    return &Handler{service: s}
}

type shortenRequest struct {
    URL string `json:"url"`
}

type shortenResponse struct {
    ShortURL string `json:"short_url"`
}

func (h *Handler) CreateShortURL(c *fiber.Ctx) error {
    req := new(shortenRequest)
    if err := c.BodyParser(req); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "invalid request")
    }

    url, err := h.service.CreateShortURL(req.URL)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }

    resp := shortenResponse{
        ShortURL: c.BaseURL() + "/" + url.ShortCode,
    }

    return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *Handler) Redirect(c *fiber.Ctx) error {
    shortCode := c.Params("id")
    url, err := h.service.GetByShortCode(shortCode)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, err.Error())
    }
    if url == nil {
        return fiber.NewError(fiber.StatusNotFound, "url not found")
    }

    return c.Redirect(url.LongURL, fiber.StatusTemporaryRedirect)
}
