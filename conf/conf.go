package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Conf struct {
	Redis Redis
	MySQL Mysql
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

var C *Conf

func init() {
	env := getMode()
	realPath, _ := filepath.Abs("./")
	viper.SetConfigName("app")
	viper.AddConfigPath(realPath + "/conf/" + env)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	C = &Conf{}
	err = viper.Unmarshal(C)
	if err != nil {
		panic(err)
	}
}

func getMode() string {
	env := os.Getenv("RUN_MODE")
	if env == "" {
		env = "dev"
	}
	return env
}
