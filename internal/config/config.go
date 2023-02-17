package config

import (
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	ENV = ".env"
)

type Config struct {
	DB *DB
}

type DB struct {
	DSN             string        `envconfig:"DATABASE_DSN"`
	MaxOpenConns    int           `envconfig:"DATABASE_MAX_OPEN_CONNS"`
	MaxIdleConns    int           `envconfig:"DATABASE_MAX_IDLE_CONNS"`
	ConnMaxLifetime time.Duration `envconfig:"DATABASE_CONN_MAX_LIFETIME"`
}

func New(ctx context.Context) (*Config, error) {
	if err := godotenv.Load(ENV); err != nil {
		return nil, fmt.Errorf("[config.New]:error loading %v: %v", ENV, err)
	}

	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("[config.New]:can't process envs: %v", err)
	}
	return &cfg, nil

}
