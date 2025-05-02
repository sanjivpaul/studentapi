package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct{
	Addr string
}

	// Env 		string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`

type Config struct{
	Env 		string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer	`yaml:"http_server"`
}

// *Config => is a return type of this function
func MustLoad() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	// if path is not find then check in flag
	if configPath == ""{
		flags:= flag.String("config", "", "path to the configuration file")

		flag.Parse()

		configPath = *flags

		// if config path is still not found then show a error
		if configPath==""{
			log.Fatal("Config path is not set")
		}
	}

	// check provide file is exist or not
	if _, err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("Config file is does not exist: %s", configPath)

	}

	// if all good then execute this
	var cfg Config

	// cleanenv is a library we are used here
	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil{
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg


}