package config

type Database struct {
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DB         string `yaml:"db"`
	Port       string `yaml:"port"`
	Host       string `yaml:"host"`
	Scheme     string `yaml:"scheme"`
	MaxConnect int    `yaml:"maxConnect"`
	Timeout    int    `yaml:"timeout"`
	SSLmod     string `yaml:"sslmod"`
}

type Config struct {
	BotToken string    `yaml:"botToken"`
	DB       *Database `yaml:"dataBase"`
}
