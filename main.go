package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

type environment struct {
	Env  string `json:"environment"`
	Port int    `json:"port"`
}

//go:embed environments
var usersDir embed.FS
var configurationFileName string

func main() {
	envPtr := flag.String("env", "development", "the build environment")
	flag.Parse()
	fmt.Println("environment:", *envPtr)
	configurationFileName = *envPtr + ".json"

	files, err := usersDir.ReadDir("environments")
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		if file.Name() == configurationFileName {
			val, err := usersDir.ReadFile("environments/" + file.Name())
			if err != nil {
				fmt.Println(err)
				continue
			}

			var environment environment
			if err := json.Unmarshal(val, &environment); err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("%+v\n", environment)
		}
	}
}
