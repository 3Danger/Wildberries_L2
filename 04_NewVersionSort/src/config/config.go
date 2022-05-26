package config

import "io"

type Config struct {
	KeysFlags *Keys
	Input     io.Reader
}

func NewConfig() *Config {
	return &Config{newKeys(), inData()}
}
