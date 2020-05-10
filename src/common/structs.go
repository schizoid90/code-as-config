package common

type (
	YamlCode struct{
		Language 	string	`yaml:"language"`
		Version		string	`yaml:"version"`
		Name		string	`yaml:"name"`
		Spec struct {
			Class 		string			`yaml:"class"`
			Functions 	[]Functions 	`yaml:"functions"`
			RunMain		string			`yaml:"runMain"`
		} `yaml:"spec"`
	}
	
	Functions struct{
		Options map[string]interface{} 	`yaml:"options,omitempty"`
		Name	string					`yaml:"name"`
		Code	map[string]interface{}	`yaml:"code"`
		Vars	map[string]interface{}	`yaml:"vars,omitempty"`
	}
)
