package library

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chrusty/tunecast/internal/configuration"
	"github.com/chrusty/tunecast/internal/storage"
	"github.com/sirupsen/logrus"
)

var (
	defaultSupportedFormats = map[string]bool{
		".flac": true,
		".mp3":  true,
		".wav":  true,
	}
)

// MediaLibrary mantains the media database:
type MediaLibrary struct {
	config           *configuration.Config
	logger           *logrus.Logger
	libraryStorage   storage.Storage
	supportedFormats map[string]bool
}

// New returns a new MediaLibrary:
func New(logger *logrus.Logger, config *configuration.Config, libraryStorage storage.Storage) (*MediaLibrary, error) {
	logger.WithField("path", config.Library.Path).Info("Preparing a new MediaLibrary ...")

	mediaLibrary := &MediaLibrary{
		config:           config,
		logger:           logger,
		libraryStorage:   libraryStorage,
		supportedFormats: defaultSupportedFormats,
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
	if l.config.Library.Scan {
		go l.scanFiles()
	}

	return nil
}

// supportedFormat determines if a file is a supported media format:
func (l *MediaLibrary) supportedFormat(path string) bool {

	ext := filepath.Ext(path)

	// Look it up:
	if _, ok := l.supportedFormats[ext]; ok {
		return true
	}

	return false
}
