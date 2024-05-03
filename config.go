package main

import (
	"encoding/json"
	"io"
	"os"
)

const configFileName = "config.json"

type Config struct {
	ModConfigSlice []ModConfig
}

func NewConfig() *Config {
	modConfigSlice := make([]ModConfig, 1)
	modConfigSlice[0] = ModConfig{
		Name: "Default",
	}
	return &Config{ModConfigSlice: modConfigSlice}
}

func (c *Config) Save() {
	new_data, _ := json.Marshal(c)
	new_string := string(new_data)

	configFile, _ := os.OpenFile(configFileName, os.O_RDONLY|os.O_CREATE, 0666)
	defer configFile.Close()

	old_data, _ := io.ReadAll(configFile)
	old_string := string(old_data)
	if new_string != old_string {
		configFile, _ := os.OpenFile(configFileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		defer configFile.Close()
		configFile.Write(new_data)
	}
}

func (c *Config) Load() {
	configFile, _ := os.Open(configFileName)
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	_ = decoder.Decode(&c)
}
