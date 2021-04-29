package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

const ERROR_UNKNOWN_FILE_TYPE = "unknown file type"

func Load(conf interface{}, file string) error {
	bts, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	ls := strings.Split(file, ".")
	if len(ls) < 2 {
		return errors.New(ERROR_UNKNOWN_FILE_TYPE)
	}
	tp := ls[len(ls)-1]
	switch tp {
	case "json":
		err := json.Unmarshal(bts, conf)
		if err != nil {
			return err
		}
		break
	case "toml":
		err := toml.Unmarshal(bts, conf)
		if err != nil {
			return err
		}
		break
	default:
		return errors.New(ERROR_UNKNOWN_FILE_TYPE)
	}
	return nil
}

func RootDir() string {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return root
}
