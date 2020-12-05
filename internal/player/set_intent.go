package player

// Pause changes the intent to be paused:
func (p *Player) Pause() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.paused = true
}

// Play changes the intent to be playing:
func (p *Player) Play() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.paused = false
}

// SetVolume changes the intended volume:
func (p *Player) SetVolume(value int32) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.volume = value
}
