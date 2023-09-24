package configs

import (
	"log/slog"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	LogLevel slog.Level `yaml:"logLevel"`

	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Port     string `yaml:"port"`
	} `yaml:"database"`

	APIServer struct {
		Port int `yaml:"port"`
	} `yaml:"apiServer"`
}

var (
	Cfg      Config
	initOnce sync.Once
)

func LoadConfig() {
	path := "./ci/"
	slog.Info("", "path", path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("read config error")
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		slog.Error("unmarshal config error")
		os.Exit(1)
	}
}

func init() {
	initOnce.Do(LoadConfig)
}
