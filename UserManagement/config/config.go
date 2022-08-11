package config

type Config struct {
	Host           string
	Port           string
	PostgresHost   string
	PostgresPort   string
	PostgresUser   string
	PostgresPass   string
	PostgresDbname string
}

func Load() Config {
	return Config{
		Host:           "localhost",
		Port:           ":2222",
		PostgresHost:   "localhost",
		PostgresPort:   "5432",
		PostgresUser:   "akbarshoh",
		PostgresPass:   "1",
		PostgresDbname: "users",
	}
}
