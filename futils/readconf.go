package futils

import (
	"encoding/json"
	"github.com/futurez/freedom/flog"
	"io/ioutil"
)

func ReadConfig(path string, cfg interface{}) error {
	if path == "" {
		path = "app.cfg"
	}
	rBuffer, err := ioutil.ReadFile(path)
	if err != nil {
		flog.Warn("[freedom] ", err.Error())
		return err
	}
	err = json.Unmarshal(rBuffer, cfg)
	if err != nil {
		flog.Warn("[freedom] ", err.Error())
		return err
	}
	return nil
}
