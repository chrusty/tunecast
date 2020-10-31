package configuration

// DatabaseConfig configures the database:
type DatabaseConfig struct {
	Disabled bool `env:"DB_DISABLED" envDefault:"true"`
}
