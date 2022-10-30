package yaag

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DocDir      string `yaml:"docDir"`
	MergeMode   string `yaml:"mergeMode"`
	SwagGeneral string `yaml:"swagGeneral"`
	YapiUrl     string `yaml:"yapiUrl"`
	YapiToken   string `yaml:"yapiToken"`
}

var YaagConfig *Config

func GetConfig() {
	if _, err := os.Stat("./yaag.yaml"); err != nil {
		fmt.Println("yaag.yaml not found, shall we create one? (y/n)")
		var input string
		fmt.Scanln(&input)
		if input != "y" {
			return
		}
		InitConfig()
	}
	// read yaag.yaml
	configYaml, err := os.ReadFile("./yaag.yaml")
	if err != nil {
		panic(err)
	}
	YaagConfig = &Config{}
	err = yaml.Unmarshal(configYaml, YaagConfig)
	if err != nil {
		panic(err)
	}
	checkConfig()
}

func InitConfig() {
	if _, err := os.Stat("./yaag.yaml"); err == nil {
		fmt.Println("yaag.yaml already exists, shall we overwrite it? (y/n)")
		var input string
		fmt.Scanln(&input)
		if input != "y" {
			return
		}
	}
	blankConfig := &Config{}
	out, _ := yaml.Marshal(blankConfig)
	err := os.WriteFile("yaag.yaml", out, 0644)
	if err != nil {
		panic(err)
	}
}

func checkConfig() {
	if YaagConfig.YapiUrl == "" || YaagConfig.YapiToken == "" {
		fmt.Println("please fill in yapiUrl and yapiToken in yaag.yaml")
		os.Exit(1)
	}
}
