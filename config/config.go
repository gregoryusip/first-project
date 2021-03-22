package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	// DBSource      string `mapstructure:"DB_SOURCE"`
	DBName        string `mapstructure:"DB_NAME"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOST"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("cannot decode into struct:", err)
	}
	return
}

func CreateConnection(path string) *sql.DB {

	config, err := LoadConfig(path)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", config.DBHost, config.DBUser, config.DBName, config.DBPassword)
	db, err := sql.Open(config.DBDriver, psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success to connect to DB!")

	return db
}
