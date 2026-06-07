package converter

import (
	"log"
	"strings"

	"github.com/itchyny/json2yaml"
)

func Convert(jsonData string) string {
	input := strings.NewReader(jsonData)

	var output strings.Builder

	err := json2yaml.Convert(&output, input)
	if err != nil {
		log.Fatalln(err)
	}

	return output.String()
}
