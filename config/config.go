package config

type Config struct {
	AllowedOrigins  string   `json:"allowedOrigins"`
	AllowedMethods  []string `json:"allowedMethods"`
	CorsCredentials bool     `json:"corsCredentials"`
	AllowedHeaders  []string `json:"allowedHeaders"`
	NameApiURL      string   `json:"nameApiUrl"`
	JokeApiURL      string   `json:"jokeApiUrl"`
	// stuff like JWT access keys
}
