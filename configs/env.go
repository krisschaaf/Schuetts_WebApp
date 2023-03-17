package configs

import (
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	MongoURIBeg string `env:"MONGOURI_BEG,required"`
	MongoURIEnd string `env:"MONGOURI_END,required"`
}

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	cfg := Config{}

	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	mongoDBPassword := os.Getenv("MONGO_DB_PASSWORD")
	mongoURI := os.ExpandEnv(cfg.MongoURIBeg + mongoDBPassword + cfg.MongoURIEnd)

	return mongoURI
}
