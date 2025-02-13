package config

import "github.com/nicholasjackson/env"

type MongoDbConfig struct {
	MongoURI string
	DBName   string
}

type AppConfig struct {
	AppURI string
}

func NewMongoDbConfig() (*MongoDbConfig, error) {
	var mongouri *string = env.String("MONGO_URI", false, "mongodb://localhost:27017", "Bind address for the Mongo server")
	var dbname *string = env.String("DB_NAME", false, "DoubtBuddy", "Bind address for the Mongo server")
	if err := env.Parse(); err != nil {
		return nil, err
	}

	return &MongoDbConfig{
		MongoURI: *mongouri,
		DBName:   *dbname,
	}, nil
}

func NewAppConfig() (*AppConfig, error) {
	var gouri *string = env.String("GO_URI", false, ":9090", "Bind address for the app server")
	if err := env.Parse(); err != nil {
		return nil, err
	}

	return &AppConfig{
		AppURI: *gouri,
	}, nil
}
