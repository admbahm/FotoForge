// Package config loads FotoForge configuration.
package config

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config contains application settings.
type Config struct {
	DatabasePath string `mapstructure:"database_path"`
	Verbose      bool   `mapstructure:"verbose"`
}

// Load reads defaults, an optional config file, environment variables, and flags.
func Load(v *viper.Viper) (Config, error) {
	v.SetDefault("database_path", filepath.Join(".fotoforge", "catalog.db"))
	v.SetEnvPrefix("FOTOFORGE")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) && !errors.Is(err, fs.ErrNotExist) {
			return Config{}, fmt.Errorf("read configuration: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("decode configuration: %w", err)
	}
	return cfg, nil
}
