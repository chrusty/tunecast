package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Recovery is a middleware that recovers from panics and logs stack-traces:
type Recovery struct {
	logger *logrus.Logger
}

// NewRecovery returns a configured recovery middleware:
func NewRecovery(logger *logrus.Logger) *Recovery {
	return &Recovery{
		logger: logger,
	}
}

// Handler returns the middleware function:
func (r *Recovery) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		// Defer a recovery function that logs the stack:
		defer func() {
			if recovery := recover(); recovery != nil {
				r.logger.
					WithField("stack", string(debug.Stack())).
					WithField("route", c.Path()).
					WithField("method", c.Request().Method).
					Errorf("Recovered from exception")

				err = fmt.Errorf("Recovered from exception")
				c.String(http.StatusInternalServerError, `{"message":"Internal Server Error"}`)
			}
		}()

		// Pass the request down the chain:
		return next(c)
	}
}
