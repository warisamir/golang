package config

import (
	"flag"
	"log"
	"os"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml:env env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:true`
	HTTPServer  `yaml:"http_server"`
}

func mustLoad() {
	var configpath string
	configpath= os.Getenv("CONFIG_PATH")


	if configpath ==""{
		flags:= flag.String("config","","path to the configuration file")
		flag.Parse()
		configpath= *flags
	}
	if configpath== "" {
		log.Fatal("Config path is not set")
	}
	if _,err:= os.Stat(configpath); os.IsNotExist(err){
		log.Fatalf("confing file not exist: %s",configpath)
	}
}