package config

import (
	"os"
)

type ConfigClick struct {
	HOST string
	PORT string
	USER string
	PASS string
}

func NewClick() *ConfigClick {

	return &ConfigClick{
		HOST: os.Getenv("CL_HOST"),
		PORT: os.Getenv("CL_PORT"),
		USER: os.Getenv("CL_USER"),
		PASS: os.Getenv("CL_PASS"),
	}
}
