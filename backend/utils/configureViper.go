package utils

import (
	"github.com/spf13/viper"
)

func ConfigureViper(confFile string) (*viper.Viper, error) {
	v := viper.New()

	// psql env bindings
	v.BindEnv("psql.user", "RDS_USERNAME")
	v.BindEnv("psql.password", "RDS_PASSWORD")
	v.BindEnv("psql.database", "RDS_DB_NAME")
	v.BindEnv("psql.host", "RDS_HOSTNAME")
	v.BindEnv("psql.port", "RDS_PORT")
	v.BindEnv("psql.ssl_mode", "RDS_SSL_MODE")
	v.BindEnv("psql.max_connections", "RDS_MAX_CONNECTIONS")
	v.BindEnv("psql.max_idle_connections", "RDS_MAX_IDLE_CONNECTIONS")
	v.BindEnv("psql.max_connection_lifetime", "RDS_MAX_CONNECTION_LIFETIME")

	// conf file
	v.SetConfigType("toml")
	v.SetConfigFile(confFile)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}
