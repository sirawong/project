package config

import "github.com/caarlos0/env"

type Config struct {
	AppPort string `env:"APP_PORT" envDefault:":8084"`
	AppUrl  string `env:"APP_URL" envDefault:"localhost:8084/reservations/checkin"`

	MongoDBEndpoint   string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName       string `env:"MONGODB_NAME" envDefault:"movie-booking"`
	MongoDBCollection string `env:"MONGODB_COLLECTION" envDefault:"reservations"`

	GRPCHost       string `env:"GRPC_HOST" envDefault:"localhost:5002"`
	GRPCAuthHost   string `env:"GRPC_AUTH_HOST" envDefault:"localhost:5000"`
	GRPCCinemaHost string `env:"GRPC_CINEMA_HOST" envDefault:"localhost:5001"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
