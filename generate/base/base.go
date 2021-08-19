package base

import (
	"fmt"
	"os"
)

func CreateFolder()  {
	err := os.Mkdir("config", 0755)
	if err != nil {
		fmt.Println("error creating folder config: " + err.Error())
		os.Exit(0)
	}

	err = os.Mkdir("services", 0755)
	if err != nil {
		fmt.Println("error creating folder services: " + err.Error())
		os.Exit(0)
	}

	err = os.Mkdir("data", 0755)
	if err != nil {
		fmt.Println("error creating folder data: " + err.Error())
		os.Exit(0)
	}
}
