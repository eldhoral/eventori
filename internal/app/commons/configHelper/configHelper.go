package configHelper

import (
	"os"

	"github.com/rs/zerolog/log"
)

func GetConfig(configName string) string {
	if val := os.Getenv(configName); val == "" {
		log.Error().Msg("err, config " + configName + " cannot be read")
	}

	return os.Getenv(configName)
}
