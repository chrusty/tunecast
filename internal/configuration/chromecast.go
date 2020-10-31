package configuration

// ChromecastConfig configures the Chromecast:
type ChromecastConfig struct {
	Address string `env:"CHROMECAST_ADDRESS" envDefault:"chromecast"`
}
