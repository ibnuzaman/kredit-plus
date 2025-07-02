package config

import "time"

type Config struct {
	Env  Env
	Port int    `env:"PORT" envDefault:"8080"`
	Host string `env:"HOST" envDefault:""`
	DB   DB     `envPrefix:"DB_"`
}

type DB struct {
	AutoMigrate        bool          `env:"AUTO_MIGRATE" envDefault:"false"`
	Name               string        `env:"NAME,required,notEmpty"`
	Host               string        `env:"HOST" envDefault:"localhost"`
	Port               int           `env:"PORT" envDefault:"5432"`
	User               string        `env:"USER,required" envDefault:""`
	Password           string        `env:"PASSWORD,required" envDefault:""`
	ConnectionIdle     time.Duration `env:"CONNECTION_IDLE" envDefault:"1m"`
	ConnectionLifetime time.Duration `env:"CONNECTION_LIFETIME" envDefault:"5m"`
	MaxIdle            int           `env:"MAX_IDLE" envDefault:"30"`
	MaxOpen            int           `env:"MAX_OPEN" envDefault:"90"`
}
