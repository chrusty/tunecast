package handler

import (
	"net/http"

	"github.com/chrusty/tunecast/api"
	"github.com/labstack/echo/v4"
)

// PutPlayerVolume - Set the intended volume (PUT /player/volume):
func (h *Handler) PutPlayerVolume(ctx echo.Context, queryParams api.PutPlayerVolumeParams) error {

	h.player.SetVolume(int32(queryParams.Value))

	// Return no content:
	return ctx.NoContent(http.StatusNoContent)
}
