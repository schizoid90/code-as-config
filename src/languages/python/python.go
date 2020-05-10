package python

import (
	"fmt"
	"os"

	c "common"
)

func Run(config c.YamlCode) {
	fmt.Println("Run python")
	outfile := fmt.Sprintf("%s.py", config.Name)
	file, err := os.Create(outfile)
	c.ChkErr(err)
	defer file.Close()
	class := config.Spec.Class
	WriteClass(file, class)
	for _, v := range config.Spec.Functions {
		WriteFunc(file, v)
	}

	fmt.Println(config.Spec.RunMain)
	if (config.Spec.RunMain != "") {
		WriteMain(file, config)
	}
}

func WriteClass(file *os.File, class string) {

	cl := fmt.Sprintf("class %s:\n", class)
	file.WriteString(cl)
}

func WriteFunc(file *os.File, fnc c.Functions){
	funcName := fmt.Sprintf("\tdef %s(self):\n", fnc.Name)
	file.WriteString(funcName)

	for k, v := range fnc.Vars {
		setVar := fmt.Sprintf("\t\t%v = %v\n", k, v)
		file.WriteString(setVar)
	}

	for c, o := range fnc.Code {
		writeCode := fmt.Sprintf("\t\t%v(%v)\n", c, o)
		file.WriteString(writeCode)
	}
}
	
func WriteMain(file *os.File, code c.YamlCode) {
	writeMain := fmt.Sprintf(`
if __name__ == '__main__':
	main = %s()
	main.%s()`, code.Spec.Class, code.Spec.RunMain)

	fmt.Println(writeMain)
	file.WriteString(writeMain)
}