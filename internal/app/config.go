package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Server struct {
		Listen string `json:"listen"`
	} `json:"server"`
	Database Database `json:"database"`
	TemplateDir string `json:"template_dir"`
}

type Database struct {
	Driver string `json:"driver"`
	Dsn    string `json:"dsn"`
	DbName string `json:"dbname"`
}

func (c *Config) LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
}

func NewConfig() Config {
	return Config{}
}
