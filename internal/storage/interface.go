package storage

import "github.com/chrusty/tunecast/api"

// Storage provides methods to work with a media library database:
type Storage interface {
	AddLibraryItem(libraryItem *api.LibraryItem) error
	List(parentPath string, sortBy string) ([]*api.LibraryItem, error)
}

// Schema (v1) - enough to work with files and folders:
// - path (string, PK)
// - uuid (UUID, unique): clients will reference this
// - folder (bool, indexed)
// - cover (string): path to the cover art (if known)
// - added (datetime, indexed): when the entry was added to the DB
