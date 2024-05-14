package config

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Config struct {
}

func DefaultConfigPath() (string, error) {
	home, err := BuildmanHome()
	if err != nil {
		return "", errors.Wrap(err, "getting buildman home")
	}
	return filepath.Join(home, "config.toml"), nil
}

func BuildmanHome() (string, error) {
	buildmanHome := os.Getenv("BUILDMAN_HOME")
	if buildmanHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", errors.Wrap(err, "getting user home")
		}
		buildmanHome = filepath.Join(home, ".buildman")
	}
	return buildmanHome, nil
}
