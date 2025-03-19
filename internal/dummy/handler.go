package dummy

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) CummulativeHandler(c *fiber.Ctx) error {
	repo := &CummulativeRepository{Db: h.db}

	// get patient id from request params
	pid := c.Params("id", "")
	if pid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":   false,
			"data": nil,
			"error": fiber.Map{
				"message": "id is required",
			},
		})
	}
	cummmulative, err := repo.GetDummyData(c.Context(), pid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"ok":   false,
			"data": nil,
			"error": fiber.Map{
				"message": err.Error(),
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ok":    true,
		"data":  cummmulative,
		"error": nil,
	})
}
