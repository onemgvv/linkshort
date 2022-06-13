package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultHTTPPort               = "8000"

	defaultPostgresPort = "5432"
	defaultPostgresHost = "localhost"

	defaultRedisPort = "6379"
	defaultRedisHost = "localhost"

	defaultLimiterRPS   = 10
	defaultLimiterBurst = 2
	defaultLimiterTTL   = 10 * time.Minute
)

type (
	Config struct {
		HTTP          HTTPConfig
		StorageConfig StorageConfig
		Limiter       LimiterConfig
		CacheTTL      time.Duration `mapstructure:"ttl"`
		Postgres      PostgresConfig
		Redis         RedisConfig
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegabytes"`
	}

	LimiterConfig struct {
		RPS   int
		Burst int
		TTL   time.Duration
	}

	StorageConfig struct {
		BaseDir string
	}

	PostgresConfig struct {
		Port     string `mapstructure:"port"`
		Host     string `mapstructure:"host"`
		Name     string `mapstructure:"dbName"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	}

	RedisConfig struct {
		Port     string `mapstructure:"port"`
		Host     string `mapstructure:"host"`
		Password string `mapstructure:"password"`
	}
)

func Init(configDir string) (*Config, error) {
	setupDefaultValues()

	if err := parseConfigFile(configDir); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)
	return &cfg, nil
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("limiter", &cfg.Limiter); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("redis", &cfg.Redis); err != nil {
		return err
	}

	return viper.UnmarshalKey("cache.ttl", &cfg.CacheTTL)
}

func setFromEnv(cfg *Config) {
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")
	cfg.Redis.Password = os.Getenv("REDIS_PASSWORD")
	cfg.StorageConfig.BaseDir = os.Getenv("STORAGE_DIR")
}

func setupDefaultValues() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.maxHeaderMegabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHTTPRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHTTPRWTimeout)

	viper.SetDefault("postgres.port", defaultPostgresPort)
	viper.SetDefault("postgres.host", defaultPostgresHost)

	viper.SetDefault("redis.port", defaultRedisPort)
	viper.SetDefault("redis.host", defaultRedisHost)

	viper.SetDefault("limiter.rps", defaultLimiterRPS)
	viper.SetDefault("limiter.burst", defaultLimiterBurst)
	viper.SetDefault("limiter.ttl", defaultLimiterTTL)
}
