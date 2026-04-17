package db

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Username      string `json:"username"`
	Bio           string `json:"bio"`
	Port          string `json:"port"`
	BroadCastPort int    `json:"broadcast_port"`
	ChatPort      int    `json:"chat_port"`
}

const (
	DefaultBroadcastPort = 8787
	DefaultChatPort      = 8989
)

func (c Config) WithDefaults() Config {
	if c.BroadCastPort == 0 {
		c.BroadCastPort = DefaultBroadcastPort
	}
	if c.ChatPort == 0 {
		c.ChatPort = DefaultChatPort
	}

	return c
}

type ConfigStore struct{}

func NewConfigStore() *ConfigStore {
	return &ConfigStore{}
}

func (s *ConfigStore) dbPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "goch", "config.json")
}

func (s *ConfigStore) Exists() bool {
	_, err := os.Stat(s.dbPath())
	return err == nil
}

func (s *ConfigStore) Save(c Config) error {
	path := s.dbPath()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func (s *ConfigStore) Load() (Config, error) {
	var c Config
	data, err := os.ReadFile(s.dbPath())
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(data, &c)
	return c, err
}
