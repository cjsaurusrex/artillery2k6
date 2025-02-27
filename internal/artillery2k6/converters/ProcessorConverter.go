package converters

import (
	"log"
	"os"
)

func ConvertProcessors(processor string) string {
	processorLines, err := os.ReadFile(processor)
	if err != nil {
		log.Fatalln(err)
	}

	return string(processorLines)
}
