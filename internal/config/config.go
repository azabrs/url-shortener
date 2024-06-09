package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)


type Config struct{
	Env string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}


type HTTPServer struct{
		Address string `yaml:"address" env-default:"localhost"`
		Timeout time.Duration `yaml:"timeout" env-default:"4s"`
		IdleTimeout time.Duration `yaml:"idle-timeout" env-default:"60s"`

	
}


func MustLoad(ConfigPath string) Config{
	if _, err := os.Stat(ConfigPath) ;err != nil {
		log.Fatalf("Config file does not exist %s", ConfigPath)
	}
	var conf Config
	if err := cleanenv.ReadConfig(ConfigPath, &conf); err != nil{
		log.Fatalf("cannot read config %s", err)
	}
	return conf
}
