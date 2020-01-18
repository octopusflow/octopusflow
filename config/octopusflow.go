package config

type DataSource struct {
	MySQL map[string]string `toml:"MySQL"`
}

type Processor struct {
	Id   string `toml:"Id"`
	Name string `toml:"Name"`
	Type string `toml:"Type"`
}

type OctopusflowConf struct {
	Processors map[string]Processor `toml:"Processors"`
}
