package configuration

// LoggingConfig configures logging:
type LoggingConfig struct {
	Level string `env:"LOGGING_LEVEL" envDefault:"INFO"`
}
