package configuration

// HTTPConfig configures the HTTP server:
type HTTPConfig struct {
	ListenAddress string `env:"HTTP_LISTEN_ADDRESS" envDefault:":8080"`
}
