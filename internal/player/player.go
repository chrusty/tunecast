package player

import (
	"sync"

	"github.com/chrusty/tunecast/api"
)

// Player maintains the "intent" of the service:
// - Playlist
// - Volume
// - Playing / paused
type Player struct {
	mutex  sync.Mutex
	paused bool
	queue  []*api.LibraryItem
}

// AddToQueue adds a libraryItem to our queue:
func (p *Player) AddToQueue(libraryItem *api.LibraryItem) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.queue = append(p.queue, libraryItem)

	return nil
}

// PlayPause changes the intent to be playing/paused:
// - Returns TRUE if paused:
func (p *Player) PlayPause() bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Switch the state:
	switch p.paused {
	case true:
		p.paused = false
	case false:
		p.paused = true
	}

	return p.paused
}
