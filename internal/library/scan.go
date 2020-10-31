package library

import (
	"os"
	"path/filepath"
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

	// Record files:
	if !info.IsDir() {
		l.logger.
			WithField("path", path).
			WithField("modified", info.ModTime().String()).
			Debug("Found a media file")
	}
	return nil
}
