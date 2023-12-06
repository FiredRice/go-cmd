package config

import (
	"encoding/json"
	"io/fs"
	"os"
)

var localConfig LocalConfig

func Save() error {
	jsonBytes, err := json.Marshal(localConfig)
	if err != nil {
		return err
	}
	err = os.WriteFile(configLocalPath, jsonBytes, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Init() error {
	content, err := os.ReadFile(configLocalPath)
	if err != nil {
		localConfig = LocalConfig{}
		return err
	}
	err = json.Unmarshal(content, &localConfig)
	if err != nil {
		localConfig = LocalConfig{}
		return err
	}
	return nil
}

func Get() LocalConfig {
	return localConfig
}

func Set(config LocalConfig) {
	localConfig = config
}
