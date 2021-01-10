package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	configuration   *Configuration
	configFileName  = "config"
	configFileExt   = ".yml"
	configType      = "yaml"
	configDirectory = "./"
	configFilePath  = filepath.Join(configDirectory, configFileName)
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	Port string `default:"8500"`
}

type DatabaseConfiguration struct {
	URI      string `default:"mongodb://localhost:27017"`
	Name     string `default:"url-shortener"`
	Username string `default:"admin"`
	Password string `default:"password"`
	Host     string `default:"localhost"`
	Port     string `default:"27017"`
	Timeout  int    `default:"10"`
}

func SetUp() (*Configuration, error) {

	initialize()

	bind()

	setDefault()

	// Read or create
	if err := readConfiguration(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

func initialize() {
	// Initialize
	viper.AddConfigPath(configDirectory)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configType)
}

func bind() {
	// Bind
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("database.uri", "DB_URI")
	viper.BindEnv("database.name", "DB_NAME")
	viper.BindEnv("database.username", "DB_USERNAME")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
}

func setDefault() {
	// Set
	viper.SetDefault("server.port", 8500)
	viper.SetDefault("database.uri", "mongodb://localhost:27017")
	viper.SetDefault("database.name", "url-shortener")
	viper.SetDefault("database.username", "admin")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 27017)
}

func readConfiguration() error {
	if err := viper.ReadInConfig(); err != nil {
		if _, err := os.Stat(configFilePath + configFileExt); os.IsNotExist(err) {
			os.Create(configFilePath + configFileExt)
		} else {
			return err
		}
	}
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}
