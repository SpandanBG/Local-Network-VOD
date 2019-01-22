package confighandler

import (
	"os"
	"bufio"
	"bytes"
	"encoding/json"
)

// GetConfigs : Read app config from config.json
func GetConfigs() (map[string]interface{}, error) {
	configData, err := readConfigFile()
	if err == nil {
		configs, err := parseConfigDataToMap(configData)
		if err == nil {
			return configs, err
		}
	}
	return nil, err
}

func readConfigFile() ([]byte, error) {
	configFile, err := os.Open("config.json")
	if err == nil {
		configReader := bufio.NewReader(configFile)
		var buffer bytes.Buffer
		parts := make([]byte, 20)
		for {
			if count, err := configReader.Read(parts); err == nil {
				buffer.Write(parts[:count])
				continue
			}
			break
		}
		return buffer.Bytes(), nil
	}
	return nil, err
}

func parseConfigDataToMap(configData []byte) (map[string]interface{}, error) {
	config := make(map[string]interface{})
	err := json.Unmarshal(configData, &config)
	return config, err
}