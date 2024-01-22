package config

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"gopkg.in/yaml.v3"
)

var (
	ENV      = "env"
	ENV_TEST = "test"
)

type Config struct {
	CurrentEnv string     `yaml:"current_env"`
	Env        ConfigData `yaml:"env"`
	EnvTest    ConfigData `yaml:"env_test"`
}

type ConfigData struct {
	DbHost     string `yaml:"db_host"`
	DbPort     string `yaml:"db_port"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
	DbName     string `yaml:"db_name"`
	DbSSLMode  string `yaml:"db_ssl_mode"`
	CookieKey  string `yaml:"cookie_key"`
}

// New returns a new Config struct
// with the values from the configuration file.
// Possible values for env are config.ENV and config.ENV_TEST.
func New(env string) *ConfigData {
	c := &Config{}
	c.ReadConfig()

	switch env {
	case ENV:
		c.CurrentEnv = ENV
		c.Env.SetRandomCookieKey()

		return &c.Env
	case ENV_TEST:
		c.CurrentEnv = ENV_TEST
		c.EnvTest.SetRandomCookieKey()

		return &c.EnvTest
	}

	return nil
}

func (c *ConfigData) SetRandomCookieKey() {
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
