package config

import "github.com/dvln/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	User   string
	Pass   string
	Host   string
	Port   int
	DBname string
}

// funcoa de ciclo de vida de golang
// para definir valores default
// viper
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5433")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.SetConfigType(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfg = &config{}

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:   viper.GetString("database.host"),
		Port:   viper.GetInt("database.port"),
		User:   viper.GetString("database.user"),
		Pass:   viper.GetString("database.pass"),
		DBname: viper.GetString("database.dbname"),
	}
	return nil
}

// initializados
func GetDBConfig() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
