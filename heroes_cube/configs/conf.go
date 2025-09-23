package configs

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DbUser string `default:"admin" split_words:"true"`
	DbPass string `default:"admin" split_words:"true"`
	DbHost string `default:"127.0.0.1" split_words:"true"`
	DbPort string `default:"3306" split_words:"true"`
	DbName string `default:"heroes" split_words:"true"`

	ServerPort string `default:"8080" split_words:"true"`
}

func LoadConfig() (*Config, error) {
	config := Config{}
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}
	return &config, nil
}
