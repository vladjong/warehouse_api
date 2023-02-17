package config

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
  ENV = ".env"
)

type Config struct {
}

func New(ctx context.Context) (*Config, error) {
  if err := godotenv.Load(ENV); err != nil {
    return nil, fmt.Errorf("[config.New]:error loading %v: %v", ENV, err)
  }

  cfg := Config{}
  if err := envconfig.Process("", &cfg); err != nil {
    return nil, fmt.Errorf("[config.New]:can't process envs: %v", ENV, err)
	}
	return &cfg, nil

}
