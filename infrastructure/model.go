package infrastructure

type app struct {
	Appname string `yaml:"name"`
	Debug   bool   `yaml:"debug"`
	Port    string `yaml:"port"`
	Service string `yaml:"service"`
	Host    string `yaml:"host"`
}

type database struct {
	Name       string `yaml:"name"`
	Connection string `yaml:"connection"`
}

type Environment struct {
	App      app                 `yaml:"app"`
	Database map[string]database `yaml:"databases"`
	path     string
}
