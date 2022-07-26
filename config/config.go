package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

// Config - config storage
type Config struct {
	DB   Database
	HTTP HTTP
	Bot  Bot
}

// Database - database config storage
type Database struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	SSLMode  string
}

// HTTP - http server config storage
type HTTP struct {
	Host      string
	Port      string
	TLSEnable bool
	CertFile  string
	KeyFile   string
}

// Bot - bot config storage
type Bot struct {
	Token string
}

// Get - read config and return as Config struct
func Get(configPath, envPrefix string) (Config, error) {
	setDefaults()

	err := read(configPath, envPrefix)
	if err != nil {
		return Config{}, err
	}

	conf := Config{
		DB: Database{
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Database: viper.GetString("db.database"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
		HTTP: HTTP{
			Host:      viper.GetString("http.host"),
			Port:      viper.GetString("http.port"),
			TLSEnable: viper.GetBool("http.tls_enable"),
			CertFile:  viper.GetString("http.cert_file"),
			KeyFile:   viper.GetString("http.key_file"),
		},
		Bot: Bot{
			Token: viper.GetString("bot.token"),
		},
	}

	return conf, nil
}

// GenerateDefault - creates config file with default config values
func GenerateDefault(configPath string) error {
	setDefaults()

	err := createConfigFile(configPath)
	if err != nil {
		return fmt.Errorf("create config failed: %v", err)
	}

	err = viper.WriteConfigAs(configPath)
	if err != nil {
		return fmt.Errorf("write to config file is failed: %v", err)
	}

	return nil
}

// read - reads config from file and environment
// configPath - filepath to config file
// envPrefix - application specific environment prefix
func read(configPath, envPrefix string) error {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix(envPrefix)
	return viper.ReadInConfig()
}

// setDefaults - set default values to viper
func setDefaults() {
	viper.SetDefault("db.username", "postgres")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.database", "database")
	viper.SetDefault("db.sslmode", "disable")
	viper.SetDefault("http.host", "localhost")
	viper.SetDefault("http.port", "8000")
	viper.SetDefault("http.tls_enable", false)
	viper.SetDefault("http.cert_file", "")
	viper.SetDefault("http.key_file", "")
	viper.SetDefault("bot.token", "")
}

// createConfigFile - creates config file
// if the file already doesn't exist, file will be created,
// otherwise file will be rewritten
func createConfigFile(configPath string) error {
	dir := filepath.Dir(configPath)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0775)
		if err != nil {
			return fmt.Errorf("create dir failed: %v", err)
		}
	}

	_, err = os.Create(configPath)
	if err != nil {
		return fmt.Errorf("create file failed: %v", err)
	}

	return nil
}
