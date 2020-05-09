package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type yamlCode struct{
	Language 	string		`yaml:"language"`
    YamlContent interface{} `yaml:"code"`
}

func main(){
	infile := flag.String("file", "", "YAML file to read")
	flag.Parse()

	codeconf, _ := filepath.Abs(*infile)
	yamlFile, err := ioutil.ReadFile(codeconf)
	chkErr(err)

	contents := readYaml(yamlFile)


    fmt.Printf("Contents: %#v\n", contents.YamlContent)
}

func readYaml(file []byte) yamlCode {
	var input yamlCode
	err := yaml.Unmarshal(file, &input)
	chkErr(err)

	return input
}

func chkErr(err error) {
	if err != nil {
        log.Fatalf("error: %v", err)
    }
}