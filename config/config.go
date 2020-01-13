package config

import (
	"fmt"
	"path/filepath"
	"sync"

	microConfig "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
)

const (
	AppPrefix = "application"
)

var rwMutex sync.RWMutex
var octopusflowConf OctopusflowConf

func Init(dir string) error {
	return initWrapper(dir, "", GetEnv())
}

func initWrapper(dir, appCommon, appEnv string) error {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	// Load config file
	err := microConfig.Load(
		// base config from application.toml file
		file.NewSource(
			file.WithPath(GetFilePath(dir, appCommon)),
		),
		// override with env file
		file.NewSource(
			file.WithPath(GetFilePath(dir, appEnv)),
		),
	)

	if err != nil {
		panic(err)
	}
	err = microConfig.Scan(&octopusflowConf)
	return err
}

func GetFilePath(dir, env string) string {
	// for dev/ci
	if dir == "" {
		dir = "../conf"
	}

	configF := AppPrefix + ".toml"
	if env != "" {
		configF = fmt.Sprintf("%s-%s.toml", AppPrefix, env)
	}
	configFp := filepath.Join(dir, configF)
	return configFp
}

func GetConfig() *OctopusflowConf {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return &octopusflowConf
}
