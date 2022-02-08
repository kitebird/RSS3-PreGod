package config

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type ServerConfig struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var Server = &ServerConfig{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("config/conf.ini")
	if err != nil {
		log.Fatalf("[error] fail to parse 'config/conf.ini': %v", err) // TODO: change to zap
	}

	mapTo("server", Server)
	Server.ReadTimeout = Server.ReadTimeout * time.Second
	Server.WriteTimeout = Server.WriteTimeout * time.Second

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err) // TODO: change to zap
	}
}
