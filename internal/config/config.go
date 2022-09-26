package config

import (
	"github.com/spf13/viper"
	"github.com/tuxoo/idler/pkg/db/mongo"
	"strings"
	"time"
)

const (
	defaultHttpPort           = "9000"
	defaultHttpRWTimeout      = 10 * time.Second
	defaultMaxHeaderMegabytes = 1
	defaultTokenTTL           = 30 * time.Minute
)

type (
	Config struct {
		HTTP  HTTPConfig
		Auth  AuthConfig
		Mongo mongo.Config
		Cache CacheConfig
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegabytes"`
	}

	AuthConfig struct {
		JWT          JWTConfig
		PasswordSalt string
	}

	JWTConfig struct {
		TokenTTL   time.Duration
		SigningKey string
	}

	CacheConfig struct {
		UserMaxSize     int
		UserExpiredTime time.Duration
	}
)

func InitConfig(path string) (*Config, error) {
	viper.AutomaticEnv()
	preDefaults()

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	if err := parseEnv(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshalConfig(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func preDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.max_header_megabytes", defaultMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHttpRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHttpRWTimeout)
	viper.SetDefault("auth.tokenTTL", defaultTokenTTL)
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0])
	viper.SetConfigName(path[1])

	return viper.ReadInConfig()
}

func parseEnv() error {
	if err := parseHttpEnv(); err != nil {
		return err
	}

	if err := parseMongoEnv(); err != nil {
		return err
	}
	return nil
}

func parseHttpEnv() error {
	if err := viper.BindEnv("http.host", "HTTP_HOST"); err != nil {
		return err
	}

	return viper.BindEnv("http.port", "HTTP_PORT")
}

func parseMongoEnv() error {
	if err := viper.BindEnv("mongo.host", "MONGO_HOST"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo.port", "MONGO_PORT"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo.db", "MONGO_INITDB_DATABASE"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo.user", "MONGO_INITDB_ROOT_USERNAME"); err != nil {
		return err
	}

	return viper.BindEnv("mongo.password", "MONGO_INITDB_ROOT_PASSWORD")
}

func unmarshalConfig(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth", &cfg.Auth.JWT); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("cache", &cfg.Cache); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mongo", &cfg.Mongo); err != nil {
		return err
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = viper.GetString("http.host")
	cfg.HTTP.Port = viper.GetString("http.port")

	cfg.Auth.PasswordSalt = viper.GetString("salt")
	cfg.Auth.JWT.SigningKey = viper.GetString("signing_key")

	cfg.Mongo.Host = viper.GetString("mongo.host")
	cfg.Mongo.Port = viper.GetString("mongo.port")
	cfg.Mongo.User = viper.GetString("mongo.user")
	cfg.Mongo.Password = viper.GetString("mongo.password")
	cfg.Mongo.DB = viper.GetString("mongo.db")
}
