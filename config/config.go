package config

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	Port string
}

type DatabaseConfiguration struct {
	URI      string
	Name     string
	Username string
	Password string
	Host     string
	Port     string
	Timeout  int
}

func SetUp() (*Configuration, error) {
	configuration := &Configuration{}
	configuration.Database = DatabaseConfiguration{
		URI:      "mongodb://localhost:27017",
		Name:     "url-shortener",
		Username: "admin",
		Password: "password",
		Host:     "localhost",
		Port:     "27017",
		Timeout:  10,
	}
	configuration.Server = ServerConfiguration{
		Port: "8500",
	}
	return configuration, nil
}
