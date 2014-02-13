package config_test

import (
	"fmt"
	"github.com/bbva-innotech/go-json-config"
)

func Example() {

	config.Load("./config.json")

	value := config.String("key-string")
	fmt.Println(value)

	values := config.StringArray("key-string-array")
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
	// 	"key-string" : "value-string"
	// }
	config.Load("./config.json")

	value := config.String("key-string")
	fmt.Println(value)
	// Output:
	// value-string
}

func ExampleStringArray() {

	// content of config.json file
	// {
	// 	"key-string-array" : [
	// 		"value-array-1",
	// 		"value-array-2",
	// 		"value-array-3"
	// 	]
	// }
	config.Load("./config.json")

	values := config.StringArray("key-string-array")
	for _, value := range values {
		fmt.Println(value)
	}
	// Output:
	// value-array-1
	// value-array-2
	// value-array-3
}

func ExampleInt() {

	// content of config.json file
	// {
	// 	"key-int" : 1
	// }
	config.Load("./config.json")

	value := config.Int("key-int")
	fmt.Println(value)
	// Output:
	// 1
}

func ExampleBool() {
	// content of config.json file
	// {
	// 	"key-bool" : true
	// }
	config.Load("./config.json")

	value := config.Bool("key-bool")
	fmt.Println(value)
	// Output:
	// true
}
