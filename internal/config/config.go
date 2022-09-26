package config

import (
	"github.com/spf13/viper"
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
		HTTPConfig HTTPConfig
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegabytes"`
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
	return parseHttpEnv()
}

func parseHttpEnv() error {
	if err := viper.BindEnv("http.host", "HTTP_HOST"); err != nil {
		return err
	}

	return viper.BindEnv("http.port", "HTTP_PORT")
}

func unmarshalConfig(cfg *Config) error {
	return viper.UnmarshalKey("http", &cfg.HTTPConfig)
}

func setFromEnv(cfg *Config) {
	cfg.HTTPConfig.Host = viper.GetString("http.host")
	cfg.HTTPConfig.Port = viper.GetString("http.port")
}
