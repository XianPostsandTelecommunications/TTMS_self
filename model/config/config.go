package config

import "time"

type AllConfig struct {
}

type Server struct {
	RunMode               string        `mapstructure:"RunMode"`
	Address               string        `mapstructure:"Address"`
	ReadTimeout           time.Duration `mapstructure:"ReadTimeout"`
	WriteTimeout          time.Duration `mapstructure:"WriteTimeout"`
	DefaultContextTimeout time.Duration `mapstructure:"DefaultContextTimeout"`
}
