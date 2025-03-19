// Routes sets up the API endpoints for the dummy module.
// It initializes a new handler with the provided database connection
// and registers the routes with the given fiber router.
//
// Parameters:
//   - app: A pointer to a fiber router where routes will be registered
//   - db: A database connection that will be passed to the handler
//
// Registered routes:
//   - GET /dummy/:id - Handles retrieving dummy data by ID
package dummy

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.Router, db *sql.DB) {
	router := *app
	handler := NewHandler(db)
	router.Get("/dummy/:id", handler.CummulativeHandler)
}
