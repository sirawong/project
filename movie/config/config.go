package config

import "github.com/caarlos0/env"

type Config struct {
	AppPort string `env:"APP_PORT" envDefault:":8000"`

	MongoDBEndpoint   string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName       string `env:"MONGODB_NAME" envDefault:"movie-booking"`
	MongoDBCollection string `env:"MONGODB_COLLECTION" envDefault:"cinema"`

	PhotoUrl string `env:"PHOTO_URL" envDefault:"http://localhost:9000/media/upload"`

	GRPCAuthHost string `env:"GRPC_AUTH_HOST" envDefault:"localhost:5000"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
