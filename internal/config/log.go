package config

type LogConfig struct {
	FileName    string `yaml:"file_name"`
	MaxSize     int    `yaml:"max_size"`
	MaxBackups  int    `yaml:"max_backups"`
	MaxKeepDays int    `yaml:"max_keep_days"`
	Compress    bool   `yaml:"compress"`
}
