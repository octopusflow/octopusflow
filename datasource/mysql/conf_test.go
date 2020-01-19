package mysql

import (
	"testing"
	"time"
)

func TestConf_MergeWithDefault(t *testing.T) {
	conf := &Conf{
		User:     "matrix",
		Password: "neo",
	}
	err := conf.MergeWithDefault()
	if err != nil {
		t.Error(err)
	}
	if conf.Host != "localhost" || conf.Timeout != 3*time.Second || conf.User != "matrix" {
		t.Errorf("conf %v merge with default failed.\n", conf)
	}
}
