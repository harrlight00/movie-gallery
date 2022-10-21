package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

// Adapted from https://onexlab-io.medium.com/golang-config-file-best-practise-d27d6a97a65a
type Configuration struct {
	JWT_KEY     string
}

func GetConfig() Configuration {
	configuration := Configuration{}

	// Get project root so tests will function as expected
	_, b, _, _ := runtime.Caller(0)

	f, err := ioutil.ReadFile(filepath.Join(filepath.Dir(b), "../..", "dev_config.json"))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(f), &configuration)
	if err != nil {
		panic(err)
	}

	return configuration
}
