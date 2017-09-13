package config

import (
	"os"
	"strings"
)

type Env int

const (
	PROD Env = iota
	STG  Env = iota
	DEV  Env = iota
)

const (
	MyHandlerPort = iota

	MyServiceHost = iota
	MyServicePort = iota
)

var prodConfig = map[int]string{
	MyHandlerPort: ":32000",

	MyServiceHost: "my-service",
	MyServicePort: ":32001",
}

var stgConfig = map[int]string{
	MyHandlerPort: ":32000",

	MyServiceHost: "my-service",
	MyServicePort: ":32001",
}

var devConfig = map[int]string{
	MyHandlerPort: ":32000",

	MyServiceHost: "my-service",
	MyServicePort: ":32001",
}

//Get gets config based on running environment.
func Get(key int) string {
	env := GetEnv()
	if env == PROD {
		return prodConfig[key]
	}
	if env == STG {
		return stgConfig[key]
	}
	return devConfig[key]
}

//Gets environment from environment variable.
func GetEnv() Env {
	env := strings.ToLower(os.Getenv("CONFIG_ENV"))
	if env == "prod" {
		return PROD
	}
	if env == "stg" {
		return STG
	}
	return DEV
}
