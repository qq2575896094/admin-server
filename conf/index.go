package conf

var Config *Conf

type Conf struct {
	Redis   *Redis   `yaml:"redis"`
	Mongodb *Mongodb `yaml:"mongodb"`
	Token   *Token   `yaml:"token"`
}
