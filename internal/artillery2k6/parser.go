package artillery2k6

import (
	"github.com/cjsaurusrex/arillery2k6/internal/artillery2k6/models"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
)

func Parse(filePath string) models.ArtilleryScript {
	input, _ := os.ReadFile(filePath)
	path, _ := filepath.Abs(filePath)
	var test = models.ArtilleryScript{Path: path}
	err := yaml.Unmarshal(input, &test)
	if err != nil {
		log.Fatal(err)
		return models.ArtilleryScript{}
	}
	return test
}
