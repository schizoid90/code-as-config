package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	
	p "languages/python"
	c "common"

	"gopkg.in/yaml.v2"
)

func main(){
	infile := flag.String("file", "", "YAML file to read")
	flag.Parse()

	codeconf, _ := filepath.Abs(*infile)
	yamlFile, err := ioutil.ReadFile(codeconf)
	c.ChkErr(err)

	contents := readYaml(yamlFile)

	if (contents.Version  != "v1"){
		fail := fmt.Sprintf("Error: version %s not recognised", contents.Version)
		panic(fail)
	}

	if (contents.Language == "python") {
		p.Run(contents)
	}
}

func readYaml(file []byte) c.YamlCode {
	var input c.YamlCode
	err := yaml.Unmarshal(file, &input)
	c.ChkErr(err)

	return input
}