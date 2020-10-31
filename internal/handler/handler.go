package handler

import (
	"github.com/chrusty/tunecast/api"
	"github.com/chrusty/tunecast/internal/library"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// Handler implements the API service:
type Handler struct {
	indentString string
	logger       *logrus.Logger
	mediaLibrary *library.MediaLibrary
}

// New returns a configured handler:
func New(logger *logrus.Logger, mediaLibrary *library.MediaLibrary) *Handler {
	return &Handler{
		indentString: "  ",
		logger:       logger,
		mediaLibrary: mediaLibrary,
	}
}

// returnError marshals and returns an error from the API spec:
func (h *Handler) returnError(ctx echo.Context, code int, details string, err error) error {
	switch {
	case code >= 500:
		h.logger.WithError(err).WithField("details", details).Errorf("Returning an error")
	case code >= 400:
		h.logger.WithError(err).WithField("details", details).Warnf("Returning an error")
	default:
		h.logger.WithError(err).WithField("details", details).Infof("Returning an error")
	}

	errorPayload := api.Error{
		Code:    int32(code),
		Message: details,
	}

	return ctx.JSONPretty(code, errorPayload, h.indentString)
}
