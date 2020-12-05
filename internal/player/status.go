package player

import (
	"github.com/chrusty/tunecast/api"
	"github.com/chrusty/tunecast/internal/utils"
)

// GetIntent returns the intended status of the player:
func (p *Player) GetIntent() *api.Status {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	volume := api.Volume(p.volume)

	intent := &api.Status{
		Volume: &volume,
	}

	// Add the track info (if we have one):
	if len(p.queue) > 0 {
		intent.Track = p.queue[0].Path
	}

	// The activity is an enum:
	switch p.paused {
	case true:
		intent.Activity = utils.String("paused")
	default:
		intent.Activity = utils.String("playing")
	}

	return intent
}

// GetStatus returns the status of the renderer:
func (p *Player) GetStatus() (*api.Status, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.renderer.Status()
}
