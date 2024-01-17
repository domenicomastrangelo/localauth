package config

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DbHost     string `yaml:"db_host"`
	DbPort     string `yaml:"db_port"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
	DbName     string `yaml:"db_name"`
	DbSSLMode  string `yaml:"db_ssl_mode"`
	CookieKey  string `yaml:"cookie_key"`
}

func New() *Config {
	c := &Config{}
	c.ReadConfig()

	c.SetRandomCookieKey()

	return c
}

func (c *Config) SetRandomCookieKey() {
	if c.CookieKey == "" {
		c.CookieKey = encryptcookie.GenerateKey()
	}
}

// ReadConfig reads the configuration file
// and stores the values in the Config.
func (c *Config) ReadConfig() {
	file, err := os.Open("/home/dome/dev/localauth_go/config.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		panic(err)
	}
}
