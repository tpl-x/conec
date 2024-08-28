package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type ServerConfig struct {
	BindPort int64      `yaml:"bind_port"`
	Log      *LogConfig `yaml:"log"`
}

func LoadDefaultConfig() (*ServerConfig, error) {
	return LoadFromFile("config.yaml")
}
func LoadFromReader(reader io.Reader) (*ServerConfig, error) {
	cfg := &ServerConfig{}
	if err := yaml.NewDecoder(reader).Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadFromFile(filePath string) (*ServerConfig, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return LoadFromReader(f)
}

func (c ServerConfig) SaveToWriter(w io.Writer) error {
	return yaml.NewEncoder(w).Encode(c)
}

func (c ServerConfig) SaveToFile(fileSavePath string) error {
	w, err := os.Create(fileSavePath)
	if err != nil {
		return err
	}
	return c.SaveToWriter(w)
}
