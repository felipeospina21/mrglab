package api

import (
	"fmt"

	"github.com/felipeospina21/mrglab/internal/config"
)

func buildURL(config *config.Config) string {
	return fmt.Sprintf("%s/api/%s", config.BaseURL, config.APIVersion)
}
