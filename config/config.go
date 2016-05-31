package config

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
)

type Auth struct {
	Token string
}
type Remote struct {
	URL string
}
type Config struct {
	Auth
	Remote
}

func getFilePath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, "gogscli.ini")
}

func ensureEndsWithSlash(str string) string {
	if len(str) == 0 {
		return str
	}
	if !strings.HasSuffix(str, "/") {
		str += "/"
	}
	return str
}

func Get() (Config, error) {
	var config Config

	iniFile, err := ini.Load(getFilePath())
	if err != nil {
		return config, err
	}

	err = iniFile.MapTo(&config)
	if err != nil {
		return config, err
	}
	config.Remote.URL = ensureEndsWithSlash(config.Remote.URL)
	return config, nil
}

func MustGet() Config {
	cfg, err := Get()
	if err != nil {
		log.Fatalf("Error getting config file: %v", err)
	}
	return cfg
}

func (c *Config) Save() error {
	iniFile := ini.Empty()
	err := ini.ReflectFrom(iniFile, &c)
	if err != nil {
		return err
	}
	return iniFile.SaveTo(getFilePath())
}
