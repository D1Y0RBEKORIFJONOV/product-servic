package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	GRPC        GRPCConfig
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	TimeOut time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config file not exist")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not exist")
	}
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("load config fail: " + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "config file path")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG")
	}
	return res
}
