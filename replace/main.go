package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {

	file := os.Args[1]
	reg := regexp.MustCompile(os.Args[2])
	repstr := os.Args[3]

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	before := reg.ReplaceAllString(string(content), repstr)
	err = ioutil.WriteFile(file, []byte(before), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
