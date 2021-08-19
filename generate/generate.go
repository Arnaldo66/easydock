package generate

import (
	"easydock/conf"
	"easydock/generate/base"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func CreateProject(language string)  {
	err := isProjectAlready()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	
	base.CreateFolder()
	createYamlFile(language)
}

func isProjectAlready() error {
	if _, err := os.Stat(conf.GetEasyDockConfFilename()); !os.IsNotExist(err) {
		return errors.New("project already exists")
	}

	return nil
}


type Element struct {
	Version string `yaml:"item"`
	Ram  string `yaml:"colour"`
}

type Services struct {
	Services map[string]Element `yaml:"services"`
}

func createYamlFile(language string)  {
	if language == "php" {
		config := Services{Services: map[string]Element{"php": {"7.4", "1024MO"},
			"postgres":{"9.6", "1024MO"}}}

		data, err := yaml.Marshal(&config)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		err = ioutil.WriteFile("config.yaml", data, 777)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
	}
}
