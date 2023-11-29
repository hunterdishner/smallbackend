package config

type Config struct {
	AllowedOrigins  string   `json:"allowedOrigins"`
	AllowedMethods  []string `json:"allowedMethods"`
	CorsCredentials bool     `json:"corsCredentials"`
	AllowedHeaders  []string `json:"allowedHeaders"`
	// stuff like JWT access keys
	// API Keys needed for other services, etc
}
