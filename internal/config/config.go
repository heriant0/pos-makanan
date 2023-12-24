package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Port string
}

type DBConfig struct {
	Driver   string
	Port     string
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	Address  string
	Password string
}

type MongoConfig struct {
	Driver   string
	UserName string
	Password string
	Host     string
	Port     string
}

type Bcrypt struct {
	HasCost int
}

type Token struct {
	Secret         string
	AccessDuration int
}

type RedirectUrl struct {
	Success string
	Failure string
}

type Payment struct {
	SecretKey     string      `yaml:"secret_key"`
	RedirectUrl   RedirectUrl `yaml:"redirect_url"`
	CallbackToken string      `yaml:"callback_token"`
}

type Ngrok struct {
	Token string `yaml:"token"`
}

type Config struct {
	App        AppConfig
	PostgresDB DBConfig
	RedisDB    RedisConfig
	MongoDB    MongoConfig
	Bcrypt     Bcrypt
	Token      Token
	Payment    Payment
	Ngrok      Ngrok
}

var Cfg *Config

func LoadConfig(fileName string) (cfg *Config) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("error:", err.Error())
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Println("error :", err.Error())
		return
	}

	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		log.Println("error :", err.Error())
		return
	}

	return
}
