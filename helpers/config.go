package helpers

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	data map[string]string
}

func NewConfig(path string) (*Config) {
	plan, ioError := ioutil.ReadFile(path)

	if ioError != nil {
		panic(ioError)
	}

	var data map[string]string
	jsonError := json.Unmarshal(plan, &data)

	if jsonError != nil {
		panic(jsonError)
	}

	return &Config{
		data: data,
	}
}

func (c *Config) Get(key string) (string) {
	return c.data[key]
}
