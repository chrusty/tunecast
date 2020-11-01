package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetLibrary - Browse the media library (GET /library):
func (h *Handler) GetLibrary(ctx echo.Context) error {

	// List the service-names from storage:
	libraryItems, err := h.mediaLibrary.List("", "")
	if err != nil {
		return h.returnError(ctx, http.StatusInternalServerError, "something broke", err)
	}

	// Return the response in the body (JSON encoded):
	return ctx.JSONPretty(http.StatusOK, libraryItems, h.indentString)
}
