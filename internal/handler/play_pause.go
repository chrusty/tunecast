package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PutPlayerPause - Pause playback (PUT /player/pause):
func (h *Handler) PutPlayerPause(ctx echo.Context) error {

	h.player.Pause()

	// Return no content:
	return ctx.NoContent(http.StatusNoContent)
}

// PutPlayerPlay - Resume playback (PUT /player/play):
func (h *Handler) PutPlayerPlay(ctx echo.Context) error {

	h.player.Play()

	// Return no content:
	return ctx.NoContent(http.StatusNoContent)
}
