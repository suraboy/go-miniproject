package internal

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server      ServerConfig      `mapstructure:"server"`
	Performance PerformanceConfig `mapstructure:"performance"`
	Database    DatabaseConfig    `mapstructure:"database"`
	Redis       RedisConfig       `mapstructure:"redis"`
}

type ServerConfig struct {
	Port          string `mapstructure:"port"`
	ReadTimeout   string `mapstructure:"read_timeout"`
	WriteTimeout  string `mapstructure:"write_timeout"`
	IdleTimeout   string `mapstructure:"idle_timeout"`
	MaxHeaderSize string `mapstructure:"max_header_size"`
}

type PerformanceConfig struct {
	MaxProcs    int    `mapstructure:"max_procs"`
	MemoryLimit string `mapstructure:"memory_limit"`
	GCPercent   int    `mapstructure:"gc_percent"`
}

type DatabaseConfig struct {
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime string `mapstructure:"conn_max_lifetime"`
}

type RedisConfig struct {
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
	DialTimeout  string `mapstructure:"dial_timeout"`
	ReadTimeout  string `mapstructure:"read_timeout"`
	WriteTimeout string `mapstructure:"write_timeout"`
}

// GetReadTimeout returns parsed read timeout
func (s *ServerConfig) GetReadTimeout() time.Duration {
	if duration, err := time.ParseDuration(s.ReadTimeout); err == nil {
		return duration
	}
	return 30 * time.Second
}

// GetWriteTimeout returns parsed write timeout
func (s *ServerConfig) GetWriteTimeout() time.Duration {
	if duration, err := time.ParseDuration(s.WriteTimeout); err == nil {
		return duration
	}
	return 30 * time.Second
}

// GetIdleTimeout returns parsed idle timeout
func (s *ServerConfig) GetIdleTimeout() time.Duration {
	if duration, err := time.ParseDuration(s.IdleTimeout); err == nil {
		return duration
	}
	return 120 * time.Second
}

func LoadConfig() (*Config, error) {
	var conf Config

	// Set file name and path
	viper.SetConfigName("config")   // Assuming config.yaml
	viper.AddConfigPath("./config") // Current directory

	// Read in the config
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Unmarshal into Config struct
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
