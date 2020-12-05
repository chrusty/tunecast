package chromecast

import (
	"fmt"

	cast "github.com/AndreasAbdi/gochromecast"
	"github.com/chrusty/tunecast/api"
	"github.com/sirupsen/logrus"
)

// Renderer implements tne Renderer interface:
type Renderer struct {
	chromecastDevice *cast.Device
	logger           *logrus.Logger
}

// New returns a configured renderer:
func New() (*Renderer, error) {
	return &Renderer{}, nil
}

// PlayLibraryItem tells the Chromecast to pause:
func (r *Renderer) PlayLibraryItem(libraryItem *api.LibraryItem) error {
	return fmt.Errorf("Unimplemented")
}

// SetPaused tells the Chromecast to pause:
func (r *Renderer) SetPaused() error {
	return fmt.Errorf("Unimplemented")
}

// SetPlaying tells the Chromecast to play:
func (r *Renderer) SetPlaying() error {
	return fmt.Errorf("Unimplemented")
}

// SetVolume tells the Chromecast to set a specific volume:
func (r *Renderer) SetVolume(percentage int) error {
	return fmt.Errorf("Unimplemented")
}

// Status returns the current status of the Chromecast:
func (r *Renderer) Status() (*api.Status, error) {
	status := &api.Status{}

	return status, nil
}
