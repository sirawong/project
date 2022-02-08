package config

import "github.com/caarlos0/env"

type Config struct {
	AppPort string `env:"APP_PORT" envDefault:":8083"`

	MongoDBEndpoint   string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName       string `env:"MONGODB_NAME" envDefault:"movie-booking"`
	MongoDBCollection string `env:"MONGODB_COLLECTION" envDefault:"movies"`

	GRPCAuthHost string `env:"GRPC_AUTH_HOST" envDefault:"localhost:5000"`

	BukgetName string `env:"BUKGET_NAME" envDefault:"asia.artifacts.movie-app-339412.appspot.com"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
