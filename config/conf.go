package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Conf struct {
	Server Server
	Redis  Redis
	MySQL  Mysql
}

type Server struct {
	Addr string
	Port int
}

type Redis struct {
	Addr string
	Pwd  string
	Db   int
}

type Mysql struct {
	Addr     string
	Username string
	Password string
}

func LoadConf() *Conf {
	env := getMode()
	viper.SetConfigName("application-" + env)
	viper.AddConfigPath("conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	conf := new(Conf)
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Panic(err)
	}
	return conf
}

func getMode() string {
	env := os.Getenv("RUN_MODE")
	if env == "" {
		env = "dev"
	}
	return env
}
