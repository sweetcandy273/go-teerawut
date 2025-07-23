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
		url = cfg.PostgreSQL.URL
	default:
		errMsg := fmt.Sprintf("error, connection url builder doesn't know the %s", stuff)
		return "", errors.New(errMsg)
	}
	return url, nil
}
