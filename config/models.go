package config

type Config struct {
	Env  Env
	Port int    `env:"PORT" envDefault:"8080"`
	Host string `env:"HOST" envDefault:""`
}
