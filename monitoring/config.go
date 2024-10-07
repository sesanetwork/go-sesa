package monitoring

// DefaultConfig is the default config for monitorings used in sesa.
type Config struct {
	Port int `toml:",omitempty"`
}

// DefaultConfig is the default config for monitorings used in sesa.
var DefaultConfig = Config{
	Port: 19090,
}
