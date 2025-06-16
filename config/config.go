// config.go
package config

var C Config

type Config struct {
	Callback CallbackConfig `yaml:"callback"`
}

type CallbackConfig struct {
	Baidu CallbackBaidu `yaml:"baidu"`
}

type CallbackBaidu struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURI  string `yaml:"redirect_uri"`
}
