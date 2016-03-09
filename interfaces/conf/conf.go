package conf

import (
	"github.com/spf13/viper"
	"fmt"
	"os"
)

type AnimalConfiguration struct {
	Name     string
	Loyal    bool
	Snickers bool
	Paws     bool
}

type Configuration struct {
	Env     string
	Animals map[string]AnimalConfiguration
}

func Get() *Configuration {
	var err error
	var c Configuration

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf/")

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal Error Reading Configuration File: %s \n", err))
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("Fatal Error Parsing Configuration into Struct: %s \n", err))
	}

	return &c;
}