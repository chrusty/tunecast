package library

import (
	"fmt"
	"os"

	"github.com/chrusty/tunecast/internal/configuration"
	"github.com/sirupsen/logrus"
)

// MediaLibrary mantains the media database:
type MediaLibrary struct {
	config *configuration.Config
	logger *logrus.Logger
}

// New returns a new MediaLibrary:
func New(logger *logrus.Logger, config *configuration.Config) (*MediaLibrary, error) {
	logger.WithField("path", config.Library.Path).Info("Preparing a new MediaLibrary")

	mediaLibrary := &MediaLibrary{
		config: config,
		logger: logger,
	}

	// Start the media library:
	if err := mediaLibrary.run(); err != nil {
		return nil, err
	}

	return mediaLibrary, nil
}

// run tells the Library to initialise:
func (l *MediaLibrary) run() error {

	// Check that the media directory exists:
	folderInfo, err := os.Stat(l.config.Library.Path)
	if err != nil {
		return err
	}

	// Check that it is a directory:
	if !folderInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", l.config.Library.Path)
	}

	// Scan for media files:
	go l.scanFiles()

	return nil
}
