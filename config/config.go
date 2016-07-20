package config

//ServerConfig is the server config struct
type ServerConfig struct {
	DBHost     string
	DBUser     string
	DBName     string
	DBSSLMode  string
	DBPassword string
	HTTPPort   int
	AdminList  []string
}
