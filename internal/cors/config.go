package cors

type Config struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

func DefaultConfig() *Config {
	return &Config{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
	}
}
