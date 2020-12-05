package player

import (
	"time"

	"github.com/chrusty/tunecast/api"
)

func (p *Player) enforceIntent(enforcementPeriod time.Duration) {
	for {
		time.Sleep(enforcementPeriod)

		// Retrieve the status of the renderer:
		rendererStatus, err := p.renderer.Status()
		if err != nil {
			continue
		}

		// Attempt to enforce the volume:
		if *rendererStatus.Volume != api.Volume(p.volume) {
			if err := p.renderer.SetVolume(p.volume); err != nil {
				p.logger.WithError(err).Error("Unable to set renderer volume")
			}
		}

		// Attempt to enforce pause:
		p.mutex.Lock()
		switch p.paused {
		case true:
			if err := p.renderer.SetPaused(); err != nil {
				p.logger.WithError(err).Error("Unable to pause renderer")
			}
		default:
			if err := p.renderer.SetPaused(); err != nil {
				p.logger.WithError(err).Error("Unable to pause renderer")
			}
		}
		p.mutex.Unlock()

	}
}
