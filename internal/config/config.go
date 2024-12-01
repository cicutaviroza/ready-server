package config

type Server struct {
	Port string `yaml:"port" env:"PORT"`
	Host string `yaml:"HOST" env:"HOST" env-default:"0.0.0.0"`
}
