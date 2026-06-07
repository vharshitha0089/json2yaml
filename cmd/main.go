package main

import (
	"fmt"

	"github.com/vharshitha0089/json2yaml/converter"
)

func main() {
	jsonData := `{"name":"Asha Rao","age":28,"active":true}`

	yamlData := converter.Convert(jsonData)

	fmt.Println(yamlData)
}
