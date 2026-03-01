package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Https struct {
	Address string
}

type Config struct {
	Env          string `yaml:"env" env-required:"true" env-default:"production" `
	Storage_path string `yaml:"storage_path" env-required:"true" `
	Https        `yaml:"http_server"`
}

func MustLoad() *Config {
	var configpath string

	configpath = os.Getenv("CONFIG_PATH")
	if configpath == ""{
		//go run ......go --config=xyz
		flags:=flag.String("config","","path to config file")
		flag.Parse()

		configpath=*flags

		if configpath==""{
			log.Fatal("Config is not set we cannot continue")
		}


	}

	if _,err:=os.Stat(configpath); os.IsNotExist(err){
		log.Fatalf("File not Found :%s",configpath)

		
	}

	var cfg Config
	err:=cleanenv.ReadConfig(configpath, &cfg)
	if err!=nil{
		log.Fatal("Not able to read config file:%s",err.Error())

	}

	return &cfg



}