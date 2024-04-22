package utils

import (
	"errors"
	"fmt"

	"github.com/sweetcandy273/go-teerawut/configs"
)

// ConnectionUrlBuilder connection url builder
func ConnectionUrlBuilder(stuff string, cfg *configs.Configs) (string, error) {
	var url string

	switch stuff {
	case "fiber":
		url = fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)
	case "postgresql":
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s",
			cfg.PostgreSQL.Host,
			cfg.PostgreSQL.Port,
			cfg.PostgreSQL.Username,
			cfg.PostgreSQL.Password,
			cfg.PostgreSQL.Database,
		)
	default:
		errMsg := fmt.Sprintf("error, connection url builder doesn't know the %s", stuff)
		return "", errors.New(errMsg)
	}
	return url, nil
}
