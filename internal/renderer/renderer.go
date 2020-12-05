package renderer

import (
	"github.com/chrusty/tunecast/api"
)

// Renderer controls audio devices (Chromecast etc):
type Renderer interface {
	PlayLibraryItem(libraryItem *api.LibraryItem) error
	SetPaused() error
	SetPlaying() error
	SetVolume(percentage int) error
	Status() (*api.Status, error)
}
