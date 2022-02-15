package config

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type HubServerConfig struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type PostgresConfig struct {
	DSN string
}

var (
	HubServer = &HubServerConfig{}
	Redis     = &RedisConfig{}
	Postgres  = &PostgresConfig{}

	cfg *ini.File
)

func Setup() error {
	var err error
	cfg, err = ini.Load("config/conf.dev.ini")

	if err != nil {
		return err
	}

	mapTo("hub-server", HubServer)
	HubServer.ReadTimeout = HubServer.ReadTimeout * time.Second
	HubServer.WriteTimeout = HubServer.WriteTimeout * time.Second

	mapTo("redis", Redis)

	mapTo("postgres", Postgres)

	return nil
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err) // TODO: change to zap
	}
}
