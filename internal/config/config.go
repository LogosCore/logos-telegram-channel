package config

import (
	"fmt"
	"os"

	coreProfile "github.com/logoscore/logos-golang-channel-core/pkg/profile"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ChannelID        string `yaml:"channel_id"`
	BotToken         string `yaml:"bot_token"`
	LogosSyncBaseURL string `yaml:"logos_sync_base_url"`
	ProfilesFile     string `yaml:"profiles_file"`
	PollTimeout      int    `yaml:"poll_timeout_seconds"`
}

func Load(path string) (Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}
	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}
	if c.ChannelID == "" {
		c.ChannelID = "telegram-main"
	}
	if c.PollTimeout <= 0 {
		c.PollTimeout = 30
	}
	if c.ProfilesFile == "" {
		c.ProfilesFile = "configs/profiles.example.yaml"
	}
	if c.BotToken == "" {
		return Config{}, fmt.Errorf("bot_token is required")
	}
	if c.LogosSyncBaseURL == "" {
		return Config{}, fmt.Errorf("logos_sync_base_url is required")
	}
	return c, nil
}

func LoadProfiles(path string) ([]coreProfile.Profile, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read profiles: %w", err)
	}
	profiles, err := coreProfile.ParseYAMLProfiles(b)
	if err != nil {
		return nil, fmt.Errorf("parse profiles: %w", err)
	}
	return profiles, nil
}
