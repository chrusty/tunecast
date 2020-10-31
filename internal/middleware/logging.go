package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Logging is a middleware that logs requests:
type Logging struct {
	logger *logrus.Logger
}

// NewLogging returns a configured logging middleware:
func NewLogging(logger *logrus.Logger) *Logging {
	return &Logging{
		logger: logger,
	}
}

// Handler returns the middleware function:
func (l *Logging) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Record the starting time:
		start := time.Now()

		// Pass the request down the chain:
		err := next(c)
		req := c.Request()
		rsp := c.Response()

		// Prepare log fields:
		logFields := logrus.Fields{
			"method":        req.Method,
			"code":          rsp.Status,
			"responseBytes": rsp.Size,
			"duration":      time.Since(start).String(),
			"route":         c.Path(),
			"path":          req.URL.Path,
			"rawQuery":      req.URL.RawQuery,
			"realClientIP":  c.RealIP(),
		}

		// Log the result (info for success, warn for failure):
		if rsp.Status < 500 {
			l.logger.WithFields(logFields).Infof("handled request")
		} else {
			l.logger.WithFields(logFields).Warnf("handled request")
		}

		return err
	}
}
