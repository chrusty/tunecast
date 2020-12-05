package handler

import (
	"net/http"

	"github.com/chrusty/tunecast/api"
	"github.com/labstack/echo/v4"
)

// GetPlayerStatus - Return the player status (GET /status/player):
func (h *Handler) GetPlayerStatus(ctx echo.Context) error {

	// Prepare a response:
	response := struct {
		Intent   *api.Status
		Renderer *api.Status
	}{
		Intent: h.player.GetIntent(),
	}

	// Get the renderer status:
	rendererStatus, err := h.player.GetStatus()
	if err != nil {
		response.Renderer = rendererStatus
	}

	// Return the response in the body (JSON encoded):
	return ctx.JSONPretty(http.StatusOK, response, h.indentString)
}

// GetRendererStatus - Return the renderer status (GET /status/renderer):
func (h *Handler) GetRendererStatus(ctx echo.Context) error {

	// Get the renderer status:
	status, err := h.player.GetStatus()
	if err != nil {
		return h.returnError(ctx, http.StatusInternalServerError, "something broke", err)
	}

	// Return the response in the body (JSON encoded):
	return ctx.JSONPretty(http.StatusOK, status, h.indentString)
}
