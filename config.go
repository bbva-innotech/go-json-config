// Package config allows you to read simple configuration values on a
// config.json file stored on your program's working directory
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var config map[string]interface{}

// Load loads json config file
func Load(filename string) {
	file, err := ioutil.ReadFile(filename)
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

// String returns a string stored on given key name
func String(name string) string {
	return config[name].(string)
}

// StringArray returns an array of strings stored on given key name
func StringArray(name string) []string {
	a := []string{}
	for _, v := range config[name].([]interface{}) {
		a = append(a, v.(string))
	}

	return a
}

// Int returns a int stored on given key
func Int(name string) int {
	i := config[name].(float64)
	return int(i)
}

// Bool returns a bool stored on given key
func Bool(name string) bool {
	return config[name].(bool)
}
