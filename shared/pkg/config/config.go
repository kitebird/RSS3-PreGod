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

var HubServer = &HubServerConfig{}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

var Redis = &RedisConfig{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config/conf.dev.ini")

	if err != nil {
		log.Fatalf("[error] fail to parse 'config/conf.ini': %v", err) // TODO: change to zap
	}

	mapTo("hub-server", HubServer)
	HubServer.ReadTimeout = HubServer.ReadTimeout * time.Second
	HubServer.WriteTimeout = HubServer.WriteTimeout * time.Second

	mapTo("redis", Redis)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err) // TODO: change to zap
	}
}
