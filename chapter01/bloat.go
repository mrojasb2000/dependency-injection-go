package main

import (
	"fmt"
	"strconv"
)

// AppendValue - get array the bytes from interface
func AppendValue(buffer []byte, in interface{}) []byte {
	var value []byte

	// convert input to []byte
	switch concrete := in.(type) {
	case []byte:
		value = concrete

	case string:
		value = []byte(concrete)
	case int64:
		value = []byte(strconv.FormatInt(concrete, 10))
	case bool:
		value = []byte(strconv.FormatBool(concrete))
	case float64:
		value = []byte(strconv.FormatFloat(concrete, 'e', 3, 64))
	}

	buffer = append(buffer, value...)
	return buffer
}

func main() {
	var data []byte
	var bytes = AppendValue(data, "hello_world")
	fmt.Println(bytes)
}
