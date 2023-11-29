package config

type Config struct {
	AllowedOrigins string `json:"allowedOrigins"`
	// stuff like JWT access keys
	// API Keys needed for other services, etc
}
