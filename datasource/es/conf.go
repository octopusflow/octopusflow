package es

import (
	"time"

	"github.com/imdario/mergo"
)

type Conf struct {
	Hosts    []string      `toml:"Hosts"`
	User     string        `toml:"User"`
	Password string        `toml:"Password"`
	Timeout  time.Duration `toml:"Timeout"` // valid time units are "ns", "us", "ms", "s"
}

var defaultConf = &Conf{
	Hosts:    []string{"localhost:9200"},
	User:     "",
	Password: "",
	Timeout:  3 * time.Second,
}

func (c *Conf) MergeWithDefault() error {
	return mergo.Merge(c, defaultConf)
}
