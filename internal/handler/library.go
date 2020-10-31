package handler

import (
	"net/http"

	"github.com/chrusty/tunecast/api"

	"github.com/labstack/echo/v4"
)

// GetLibrary - Browse the media library (GET /library):
func (h *Handler) GetLibrary(ctx echo.Context) error {

	// List the service-names from storage:
	response := &[]api.LibraryItem{}

	// Return the response in the body (JSON encoded):
	return ctx.JSONPretty(http.StatusOK, response, h.indentString)
}
