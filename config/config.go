package config

import (
	"errors"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var Set = wire.NewSet(
	NewConfig,
)

type Configuration struct {
	Server ServerConfig
	Logger Logger
	Kafka  KafkaConfig
}

// ServerConfig struct
type ServerConfig struct {
	Name                string
	AppVersion          string
	Port                string
	BaseURI             string
	Mode                string
	Prefork             bool
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
	SSL                 bool
	CtxDefaultTimeout   time.Duration
	CSRF                bool
	Debug               bool
	GrRunningThreshold  int //  threshold for goroutines are running (which could indicate a resource leak).
	GcPauseThreshold    int //  threshold garbage collection pause exceeds. (Millisecond)
	CacheDeploymentType int
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// KafkaConfig struct
type KafkaConfig struct {
	Brokers            []string
	TopicPrefix        string
	DefaultPartitions  int32
	DefaultReplication int16
}

var configSettings = Configuration{}

// NewConfig Load config file from given path
func NewConfig() (*Configuration, error) {
	path := os.Getenv("cfgPath")
	if path == "" {
		path = getDefaultConfig()
	}

	v := viper.New()

	v.SetConfigName(path)
	v.SetConfigType("yaml") // Specify the config file type (e.g., yaml, json)
	v.AddConfigPath(".")    // Look for the config file in the current directory
	v.AddConfigPath("/app") // Look for the config file in the Docker container path
	v.AddConfigPath("/config")

	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	err := v.Unmarshal(&configSettings)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &configSettings, nil
}

// Get config path for local or docker
func getDefaultConfig() string {
	return "config/config-local"
}
