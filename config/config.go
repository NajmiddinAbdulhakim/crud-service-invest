package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string


	RPCPort string
}

func Load() Config {
	c := Config{}

	c.PostgresHost = cast.ToString(look("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(look("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(look("POSTGRES_DATABASE", "iman2"))
	c.PostgresUser = cast.ToString(look("POSTGRES_USER", "najmiddin"))
	c.PostgresPassword = cast.ToString(look("POSTGRES_PASSWORD", "1234"))

	c.RPCPort = cast.ToString(look("RPC_PORT", ":9000"))

	return c
}

func look(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
