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

var (
	outfile *string
)

func main(){
	infile := flag.String("file", "", "YAML file to read")
	outfile = flag.String("out", "output.txt", "File to create")
	flag.Parse()

	codeconf, _ := filepath.Abs(*infile)
	yamlFile, err := ioutil.ReadFile(codeconf)
	c.ChkErr(err)

	contents := readYaml(yamlFile)


	fmt.Printf("--- t:\n%v\n\n", contents)
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