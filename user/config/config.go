package config

import "github.com/caarlos0/env"

type Config struct {
	AppPort string `env:"APP_PORT" envDefault:":8080"`

	MongoDBEndpoint   string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName       string `env:"MONGODB_NAME" envDefault:"movie-dev"`
	MongoDBCollection string `env:"MONGODB_COLLECTION" envDefault:"user"`

	PhotoUrl string `env:"PHOTO_URL" envDefault:"http://localhost:9000/media/upload"`

	GRPCHost string `env:"GRPC_HOST" envDefault:"localhost:5000"`

	JWTSecret string `env:"JWT_SECRET" envDefault:"secret"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
