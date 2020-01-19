package mysql

import (
	"fmt"
	"time"

	"github.com/imdario/mergo"
)

type Conf struct {
	Host         string        `toml:"Host"`
	Port         int           `toml:"Port"`
	DBName       string        `toml:"DBName"`
	DBCharset    string        `toml:"DBCharset"`
	User         string        `toml:"User"`
	Password     string        `toml:"Password"`
	Timeout      time.Duration `toml:"Timeout"` // Dial timeout, valid time units are "ns", "us", "ms", "s"
	WriteTimeout time.Duration `toml:"WriteTimeout"`
	ReadTimeout  time.Duration `toml:"ReadTimeout"`
	MaxIdleConns int           `toml:"MaxIdleConns"`
	MaxOpenConns int           `toml:"MaxOpenConns"`
}

var defaultConf = &Conf{
	Host:         "localhost",
	Port:         3306,
	DBCharset:    "utf8mb4",
	Timeout:      3 * time.Second,
	WriteTimeout: 3 * time.Second,
	ReadTimeout:  3 * time.Second,
	MaxIdleConns: 8,
	MaxOpenConns: 8,
}

func (c *Conf) MergeWithDefault() error {
	return mergo.Merge(c, defaultConf)
}

func (c *Conf) BuildArgs() string {
	prefix := fmt.Sprintf("%s:%s@(%s)/%s?", c.User, c.Password, c.Host, c.DBName)
	suffix := fmt.Sprintf("charset=%s&parseTime=True&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s",
		c.DBCharset, c.Timeout, c.ReadTimeout, c.WriteTimeout)
	return prefix + suffix
}
