package config

import (
	"encoding/json"
	"os"
	"strings"
)

func Load(fileName string) (Configuration, error) {
	var data []byte
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(strings.NewReader(string(data)))
	m := make(map[string]interface{})

	err = decoder.Decode(&m)
	if err != nil {
		return nil, err
	}

	return &DefaultConfig{configData: m}, nil
}
