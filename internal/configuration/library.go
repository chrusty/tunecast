package configuration

// LibraryConfig configures the media library:
type LibraryConfig struct {
	Path string `env:"LIBRARY_PATH" envDefault:"/media"`
}
