package es

import (
	"testing"
	"time"
)

func TestConf_MergeWithDefault(t *testing.T) {
	testHost := "https://es.test.com:9200"
	conf := &Conf{
		Hosts:    []string{testHost},
		User:     "matrix",
		Password: "neo",
	}
	err := conf.MergeWithDefault()
	if err != nil {
		t.Error(err)
	}
	if conf.Hosts[0] != testHost || conf.Timeout != 3*time.Second || conf.User != "matrix" {
		t.Errorf("conf %v merge with default failed.\n", conf)
	}
}
