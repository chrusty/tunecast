package chromecast

import (
	"fmt"
	"net"
	"time"

	cast "github.com/AndreasAbdi/gochromecast"
	"github.com/AndreasAbdi/gochromecast/controllers/receiver"
	"github.com/chrusty/tunecast/api"
	"github.com/chrusty/tunecast/internal/configuration"
	"github.com/chrusty/tunecast/internal/utils"
	"github.com/sirupsen/logrus"
)

// Renderer implements tne Renderer interface:
type Renderer struct {
	chromecastDevice *cast.Device
	controlTimeout   time.Duration
	logger           *logrus.Logger
}

// New returns a configured renderer:
func New(logger *logrus.Logger, config *configuration.Config) (*Renderer, error) {
	logger.WithField("address", config.Chromecast.Address).WithField("port", config.Chromecast.Port).Info("Preparing a new Chromecast renderer ...")

	// Make a new Chromecast device:
	chromeCastIP := net.ParseIP(config.Chromecast.Address)
	chromecastDevice, err := cast.NewDevice(chromeCastIP, config.Chromecast.Port)
	if err != nil {
		return nil, err
	}

	return &Renderer{
		chromecastDevice: &chromecastDevice,
		controlTimeout:   time.Second,
		logger:           logger,
	}, nil
}

// PlayLibraryItem tells the Chromecast to pause:
func (r *Renderer) PlayLibraryItem(libraryItem *api.LibraryItem) error {
	return fmt.Errorf("Unimplemented")
}

// SetPaused tells the Chromecast to pause:
func (r *Renderer) SetPaused() error {
	_, err := r.chromecastDevice.MediaController.Pause(r.controlTimeout)
	return err
}

// SetPlaying tells the Chromecast to play:
func (r *Renderer) SetPlaying() error {
	_, err := r.chromecastDevice.MediaController.Play(r.controlTimeout)
	return err
}

// SetVolume tells the Chromecast to set a specific volume:
func (r *Renderer) SetVolume(percentage int32) error {
	newVolume := &receiver.Volume{
		Level: utils.Float64(float64(percentage)),
	}

	_, err := r.chromecastDevice.ReceiverController.SetVolume(newVolume, r.controlTimeout)
	return err
}

// Status returns the current status of the Chromecast:
func (r *Renderer) Status() (*api.Status, error) {

	mediaStatus := r.chromecastDevice.GetMediaStatus(r.controlTimeout)
	if len(mediaStatus) == 0 {
		return nil, fmt.Errorf("Nothing playing")
	}

	status := &api.Status{}

	return status, nil
}
