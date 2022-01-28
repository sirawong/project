package config

import "github.com/caarlos0/env"

type Config struct {
	AppPort string `env:"APP_PORT" envDefault:":8080"`
	AppUrl  string `env:"APP_URL" envDefault:":8080/checkin"`

	MongoDBEndpoint   string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName       string `env:"MONGODB_NAME" envDefault:"movie-booking"`
	MongoDBCollection string `env:"MONGODB_COLLECTION" envDefault:"reservation"`

	GRPCHost       string `env:"GRPC_HOST" envDefault:"localhost:5002"`
	GRPCAuthHost   string `env:"GRPC_AUTH_HOST" envDefault:"localhost:5000"`
	GRPCCinemaHost string `env:"GRPC_CINEMA_HOST" envDefault:"localhost:5001"`

	SendgridAPIKey string `env:"SENDGRID_API_KEY" envDefault:"..."`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
