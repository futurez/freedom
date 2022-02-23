package futils

import (
	"encoding/json"
	"github.com/futurez/freedom/flog"
	"io/ioutil"
)

func ReadConfig(path string, cfg interface{}) {
	if path == "" {
		path = "app.cfg"
	}
	rBuffer, err := ioutil.ReadFile(path)
	if err != nil {
		flog.Fatal(err)
	}
	err = json.Unmarshal(rBuffer, cfg)
	if err != nil {
		flog.Fatal(err)
	}
}
