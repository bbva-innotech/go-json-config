package config_test

import (
	"fmt"
	"github.com/bbva-innotech/go-json-config"
)

func Example() {

	// content of config.json file
	// {
	// 	"key-string" : "value-string",
	// 	"key-array" : [
	// 		"value-array-1",
	// 		"value-array-2",
	// 		"value-array-3"
	// 	]
	// }

	value := config.String("key-string")
	fmt.Println(value)

	values := config.StringArray("key-array")
	for _, value := range values {
		fmt.Println(value)
	}
	// Output:
	// value-string
	// value-array-1
	// value-array-2
	// value-array-3
}

func ExampleString() {

	// content of config.json file
	// {
	// 	"key-string" : "value-string",
	// 	"key-array" : [
	// 		"value-array-1",
	// 		"value-array-2",
	// 		"value-array-3"
	// 	]
	// }

	value := config.String("key-string")
	fmt.Println(value)
	// Output:
	// value-string
}

func ExampleStringArray() {

	// content of config.json file
	// {
	// 	"key-string" : "value-string",
	// 	"key-array" : [
	// 		"value-array-1",
	// 		"value-array-2",
	// 		"value-array-3"
	// 	]
	// }

	values := config.StringArray("key-array")
	for _, value := range values {
		fmt.Println(value)
	}
	// Output:
	// value-array-1
	// value-array-2
	// value-array-3
}
