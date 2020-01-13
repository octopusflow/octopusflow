package config

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestGetFilePath(t *testing.T) {
	appCommon := GetFilePath("", "")
	appCommonAbs, _ := filepath.Abs(appCommon)
	target := "octopusflow/conf/application.toml"
	if !strings.HasSuffix(appCommonAbs, target) {
		t.Errorf("app common %s != %s\n", appCommonAbs, target)
	}
}

func TestInit(t *testing.T) {
	err := Init("")
	if err != nil {
		t.Error(err)
	}
}

func TestGetConfig(t *testing.T) {
	typeCi := "ci"
	err := initWrapper("", "ci-common", "ci")
	if err != nil {
		t.Error(err)
	}
	conf := GetConfig()
	ciProc := conf.Processors["ci"]
	idTarget := "ci"
	if ciProc.Id != idTarget {
		t.Errorf("ci processor id %s != %s", ciProc.Id, idTarget)
		return
	}
	nameTarget := "ci"
	if ciProc.Name != nameTarget {
		t.Errorf("ci processor name %s != %s", ciProc.Name, nameTarget)
		return
	}
	if ciProc.Type != typeCi {
		t.Errorf("ci processor type %s != %s", ciProc.Type, typeCi)
		return
	}
}
