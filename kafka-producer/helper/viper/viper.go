package viper

import (
	"strings"

	"github.com/spf13/viper"
)

// Config ...
type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	Init()
}

type config struct{}

// New ...
func NewViper() Config {
	v := &config{}
	v.Init()
	return v
}

// Init ...
func (v *config) Init() {
	viper.SetEnvPrefix(`test`)
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(`config.json`)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// GetString ...
func (v *config) GetString(key string) string {
	return viper.GetString(key)
}

// GetInt ...
func (v *config) GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool ...
func (v *config) GetBool(key string) bool {
	return viper.GetBool(key)
}
