package conf

type Config struct {
	Mysql Mysql `yaml:"mysql"`
}

type Mysql struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
