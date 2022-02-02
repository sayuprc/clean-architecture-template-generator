package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Directory struct {
	Root string `json:"root"`
}

type Namespace struct {
	Entity             string `json:"entity"`
	Repository         string `json:"repository"`
	InMemoryRepository string `json:"inmemoryrepository"`
	UseCase            string `json:"usecase"`
	Interactor         string `json:"interactor"`
	MockInteractor     string `json:"mockinteractor"`
}

type Config struct {
	Directory Directory `json:"directory"`
	Namespace Namespace `json:"namespace"`
}

func Load() (*Config, error) {
	viper.SetConfigName(".cagt")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return &cfg, nil
}
