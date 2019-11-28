package apiconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// APIConfig is
type APIConfig struct {
	Server struct {
		Addr string `json:"addr"`
	} `json:"server"`
}

var instance APIConfig

// Start is
func Start(env string) (err error) {
	executedPath, err := os.Executable()
	if err != nil {
		return
	}

	path := fmt.Sprintf("../../config/config.%s.json", env)
	filepath.Join(executedPath)
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&instance)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	return
}

// Get is
func Get() *APIConfig {
	return &instance
}
