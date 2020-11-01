package library

import "github.com/chrusty/tunecast/api"

// List just passes the request on to the storage implementation:
func (l *MediaLibrary) List(parentPath string, sortBy string) ([]*api.LibraryItem, error) {
	return l.libraryStorage.List(parentPath, sortBy)
}
