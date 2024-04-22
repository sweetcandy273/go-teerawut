package configs

// Configs configs
type Configs struct {
	PostgreSQL PostgreSQL
	App        Fiber
}

// Fiber fiber
type Fiber struct {
	Host string
	Port string
}

// PostgreSQL postgresql
type PostgreSQL struct {
	Host     string
	Port     string
	Protocol string
	Username string
	Password string
	Database string
	SSLMode  string
}
