package config

import "github.com/arrowwhi/go-utils/grpcserver/grpc_config"

type Config struct {
	Config grpc_config.Config

	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"` // Уровень логирования

	PG Postgres `envconfig:"POSTGRES" required:"true"`
}

type Postgres struct {
	User     string `envconfig:"USER" required:"true"`     // Логин для подключения
	Password string `envconfig:"PASSWORD" required:"true"` // Пароль для подключения
	Database string `envconfig:"DB" required:"true"`       // Название базы данных
	Host     string `envconfig:"HOST" required:"true"`     // Хост базы данных
	Port     string `envconfig:"PORT" required:"true"`     // Порт базы данных
}
