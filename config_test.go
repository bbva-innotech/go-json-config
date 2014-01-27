package config_test

import (
	"fmt"
	"github.com/bbva-innotech/go-json-config"
)

func ExampleString() {

	// content of config.json file
	// {
	// 	"key-string" : "value",
	// 	"key-array" : [
	// 		"value1",
	// 		"value2",
	// 		"value3"
	// 	]
	// }

	value := config.String("key-string")
	fmt.Println(value)
	// Output: value1
}

func ExampleStringArray() {

	// content of config.json file
	// {
	// 	"key-string" : "value",
	// 	"key-array" : [
	// 		"value1",
	// 		"value2",
	// 		"value3"
	// 	]
	// }

	values := config.StringArray("key-array")
	for _, value := range values {
		fmt.Println(value)
	}
	// Output:
	// value1
	// value2
	// value3
}
