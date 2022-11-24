package api

import (
	"os"

	yaml "gopkg.in/yaml.v3"
	utils "wgen/mod/utils"
)

type YAPI struct {
	RelativeUrl string `yaml:"relativeUrl"`
	Method      string `yaml:"method"`
}

type YAPISpec struct {
	Name    string          `yaml:"name"`
	BaseUrl string          `yaml:"baseUrl"`
	APIs    map[string]YAPI `yaml:"api"`
}

func ParseYamlAPISpecFile(filename string) YAPISpec {
	file, err := os.ReadFile(filename)
	if err != nil {
		utils.Kill(err)
	}

	var APISpec YAPISpec
	err = yaml.Unmarshal(file, &APISpec)
	if err != nil {
		utils.Kill(err)
	}

	return APISpec
}
