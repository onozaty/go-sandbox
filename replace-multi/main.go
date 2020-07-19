package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// Setting 置換設定
type Setting struct {
	File        string `json:"file"`
	Regex       string `json:"regex"`
	Replacement string `json:"replacement"`
}

func main() {

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var settings []Setting
	if err := json.Unmarshal(content, &settings); err != nil {
		log.Fatal(err)
	}

	for _, setting := range settings {

		content, err := ioutil.ReadFile(setting.File)
		if err != nil {
			log.Fatal(err)
		}

		regex := regexp.MustCompile(setting.Regex)

		before := regex.ReplaceAllString(string(content), setting.Replacement)
		err = ioutil.WriteFile(setting.File, []byte(before), 0666)
		if err != nil {
			log.Fatal(err)
		}
	}
}
