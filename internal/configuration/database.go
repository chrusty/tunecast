package configuration

import "fmt"

const (
	sqlDBFileName = "tunecast.db"
)

// DatabaseConfig configures the database:
type DatabaseConfig struct {
	SQLitePath string `env:"DB_SQLITE_PATH" envDefault:"/library"`
}

// SQLiteDBPath returns the full path to the SQLite DB file:
func (d *DatabaseConfig) SQLiteDBPath() string {
	return fmt.Sprintf("%s/%s", d.SQLitePath, sqlDBFileName)
}
