package config

import (
	"flag"
	"time"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type AgentConfig struct {
	Address        string        `env:"ADDRESS"`
	PollInterval   time.Duration `env:"POLL_INTERVAL"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
	Key            string        `env:"KEY"`
}

type StorageConfig struct {
	StoreFile     string        `env:"STORE_FILE"`
	StoreInterval time.Duration `env:"STORE_INTERVAL"`
	Restore       bool          `env:"RESTORE"`
	DatabaseDNS   string        `env:"DATABASE_DSN"`
}

type ServerConfig struct {
	Address       string `env:"ADDRESS"`
	Key           string `env:"KEY"`
	StorageConfig StorageConfig
}

func GetAgentConfig() AgentConfig {
	var config AgentConfig

	flag.StringVar(&config.Address, "a", "127.0.0.1:8080", "Server address")
	flag.DurationVar(&config.PollInterval, "p", time.Second*2, "Poll interval, sec")
	flag.DurationVar(&config.ReportInterval, "r", time.Second*10, "Report interval, sec")
	flag.StringVar(&config.Key, "k", "", "HMAC key")
	flag.Parse()

	if err := env.Parse(&config); err != nil {
		log.Fatalf("Failed to parse agent config, %v", err)
	}

	return config
}

func GetServerConfig() ServerConfig {
	var config ServerConfig

	flag.StringVar(&config.Address, "a", "127.0.0.1:8080", "Server Address")
	flag.StringVar(&config.Key, "k", "", "HMAC key")
	flag.StringVar(&config.StorageConfig.StoreFile, "f", "/tmp/devops-metrics-db.json", "Store File")
	flag.DurationVar(&config.StorageConfig.StoreInterval, "i", time.Second*300, "Store Interval")
	flag.BoolVar(&config.StorageConfig.Restore, "r", true, "Restore After Start")
	flag.StringVar(&config.StorageConfig.DatabaseDNS, "d", "", "Database DNS")
	flag.Parse()

	if err := env.Parse(&config); err != nil {
		log.Fatalf("Failed to parse server config, %v", err)
	}

	return config
}
