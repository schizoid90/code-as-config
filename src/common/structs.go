package common

type YamlCode struct{
	// Arbitrarily read YAML
	Language 	string		`yaml:"language"`
    YamlContent interface{} `yaml:"content"`
}