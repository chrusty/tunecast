package library

import (
	"os"
	"path/filepath"

	"github.com/chrusty/tunecast/api"
	"github.com/chrusty/tunecast/internal/utils"
)

// scanFiles walks the media directory:
func (l *MediaLibrary) scanFiles() {
	if err := filepath.Walk(l.config.Library.Path, l.fileWalkFunc); err != nil {
		l.logger.WithError(err).Fatalf("Error scanning media library files")
	}
}

// fileWalkFunc is run for every file we find:
func (l *MediaLibrary) fileWalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	relativePath := path[len(l.config.Library.Path):]

	// Record folders:
	if info.IsDir() {
		l.logger.
			WithField("path", path).
			WithField("modified", info.ModTime().String()).
			Debug("Found a library folder")

		libraryItem := &api.LibraryItem{
			Cover:    utils.String(""),
			IsFolder: utils.Bool(true),
			Path:     utils.String(relativePath),
		}

		return l.libraryStorage.AddLibraryItem(libraryItem)
	}

	// Record media files:
	if l.supportedFormat(path) {
		l.logger.
			WithField("path", path).
			WithField("modified", info.ModTime().String()).
			Debug("Found a library file")

		libraryItem := &api.LibraryItem{
			Cover:    utils.String(""),
			IsFolder: utils.Bool(false),
			Path:     utils.String(relativePath),
		}

		return l.libraryStorage.AddLibraryItem(libraryItem)
	}

	// If we get this far then we found a file with an unsupported format:
	l.logger.
		WithField("path", path).
		Warn("Unsupported format")

	return nil
}
