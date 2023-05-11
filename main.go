package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Generated from YAML sample with https://zhwt.github.io/yaml-to-go/
// Reference: https://stackoverflow.com/questions/28682439/go-parse-yaml-file
type SloDefinition struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name   string `yaml:"name"`
		Labels struct {
			ServiceName string `yaml:"service_name"`
			FeatureName string `yaml:"feature_name"`
			SloName     string `yaml:"slo_name"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
	Spec struct {
		Description           string   `yaml:"description"`
		Backend               string   `yaml:"backend"`
		Method                string   `yaml:"method"`
		Exporters             []string `yaml:"exporters"`
		ServiceLevelIndicator struct {
			FilterGood  string `yaml:"filter_good"`
			FilterValid string `yaml:"filter_valid"`
		} `yaml:"service_level_indicator"`
		Goal float64 `yaml:"goal"`
	} `yaml:"spec"`
}

func main() {
	yamlFile, _ := filepath.Abs("./slo_gae_app_availability.yaml")
    yamlContent, err := os.ReadFile(yamlFile)
	if err != nil {
        panic(err)
    }

	var sloDefinition SloDefinition
	err = yaml.Unmarshal(yamlContent, &sloDefinition)
	if err != nil {
		panic(err)
	}

	fmt.Println(sloDefinition.Metadata.Name)
	fmt.Println(sloDefinition.Spec.Backend)
	fmt.Println(sloDefinition.Spec.ServiceLevelIndicator.FilterGood)
}
