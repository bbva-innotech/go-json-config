// Package config allows you to read simple configuration values on a
// config.json file stored on your programs working directory
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var config map[string]interface{}

func init() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(string(file)), &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v\n", err)
		os.Exit(1)
	}
}

// Returns a string stored on given key name
func String(name string) string {
	return config[name].(string)
}

// Returns an array of strings stored on given key name
func StringArray(name string) []string {
	a := []string{}
	for _, v := range config[name].([]interface{}) {
		a = append(a, v.(string))
	}

	return a
}
