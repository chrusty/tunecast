package main

import (
	"net/http"

	"github.com/chrusty/tunecast/api"
	"github.com/chrusty/tunecast/internal/configuration"
	"github.com/chrusty/tunecast/internal/handler"
	"github.com/chrusty/tunecast/internal/library"
	"github.com/chrusty/tunecast/internal/middleware"

	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {

	// Make a Logrus logger:
	logger := logrus.New()

	// Load config:
	config, err := configuration.Load()
	if err != nil {
		logger.WithError(err).Fatal("Unable to load config")
	}

	// Set the logging level specified in the config:
	loggingLevel, err := logrus.ParseLevel(config.Logging.Level)
	if err != nil {
		logger.WithError(err).Fatal("Invalid log level")
	}
	logger.SetLevel(loggingLevel)

	// Prepare a media library:
	mediaLibrary, err := library.New(logger, config)
	if err != nil {
		logger.WithError(err).Fatal("Unable to prepare a media library")
	}

	// Prepare an API handler:
	apiHandler := handler.New(logger, mediaLibrary)

	// Load the OpenAPI spec:
	openAPISpec, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	// Clear out the servers array in the API spec, that skips validating that server names match:
	openAPISpec.Servers = nil

	// Get a basic echo router:
	echoRouter := echo.New()

	// Hide the banner in the logs:
	echoRouter.HideBanner = true

	// Recovery middleware:
	echoRouter.Use(middleware.NewRecovery(logger).Handler)

	// GZIP middleware:
	echoRouter.Use(echoMiddleware.Gzip())

	// Logging middleware:
	echoRouter.Use(middleware.NewLogging(logger).Handler)

	// CORS middleware:
	echoRouter.Use(echoMiddleware.CORS())

	// OpenAPI request validation middleware:
	echoRouter.Use(oapiMiddleware.OapiRequestValidatorWithOptions(openAPISpec, &oapiMiddleware.Options{}))

	// Add the handler:
	api.RegisterHandlers(echoRouter, apiHandler)

	// Serve with Mux (allows us to host static content and the API together):
	muxRouter := mux.NewRouter()
	muxRouter.PathPrefix("/api/v1").Handler(http.StripPrefix("/api/v1", echoRouter))
	logger.Infof("Listening on %s ...", config.HTTP.ListenAddress)
	logger.WithError(http.ListenAndServe(config.HTTP.ListenAddress, muxRouter)).Fatal("Stopped listening")
}
