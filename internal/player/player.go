package player

import (
	"sync"
	"time"

	"github.com/chrusty/tunecast/api"
	"github.com/chrusty/tunecast/internal/renderer"
	"github.com/sirupsen/logrus"
)

// Player maintains the "intent" of the service:
// - Playlist
// - Volume
// - Playing / paused
type Player struct {
	logger   *logrus.Logger
	mutex    sync.Mutex
	paused   bool
	queue    []*api.LibraryItem
	renderer renderer.Renderer
	volume   int32
}

// New returns a configured Player:
func New(logger *logrus.Logger, renderer renderer.Renderer) *Player {
	logger.Info("Preparing a new player ...")

	newPlayer := &Player{
		logger:   logger,
		renderer: renderer,
		volume:   75,
	}

	go newPlayer.enforceIntent(time.Second)

	return newPlayer
}

// AddToQueue adds a libraryItem to our queue:
func (p *Player) AddToQueue(libraryItem *api.LibraryItem) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.queue = append(p.queue, libraryItem)

	return nil
}
