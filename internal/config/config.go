package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server `mapstructure:"server"`
}

type Server struct {
	Addr string `mapstructure:"addr"`
}

type Service struct {
	cfg Config
}

func New(path string) (*Service, error) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("parse config %s: %w", path, err)
	}
	return &Service{cfg: c}, nil
}

func (s *Service) Get() Config { return s.cfg }
