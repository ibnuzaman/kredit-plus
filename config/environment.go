package config

import (
	"os"
	"strings"
)

type Env string

const (
	EnvDevelopment Env = "development"
	EnvProduction  Env = "production"
	EnvStaging     Env = "staging"
)

func (e Env) String() string {
	return string(e)
}

func (e Env) IsDevelopment() bool {
	return e == EnvDevelopment
}

func (e Env) IsProduction() bool {
	return e == EnvProduction
}

func (e Env) IsStaging() bool {
	return e == EnvStaging
}

func (e Env) IsValid() bool {
	switch e {
	case EnvDevelopment, EnvProduction, EnvStaging:
		return true
	}
	return false
}

func (e Env) FromEnv(key string) Env {
	defaultValue := EnvDevelopment
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	env := Env(strings.ToLower(strings.TrimSpace(value)))
	if !env.IsValid() {
		return defaultValue
	}

	return env
}
