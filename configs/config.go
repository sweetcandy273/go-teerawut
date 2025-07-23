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
	URL string
}
