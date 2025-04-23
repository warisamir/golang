package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}
//  env-default:"production"
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:true`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config{
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
	var cfg Config
	err:= cleanenv.ReadConfig(configpath, &cfg)
	if err !=nil{
		log.Fatalf("cannot read config file : %s",err.Error())
	}
	return &cfg;
}