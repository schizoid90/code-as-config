package python

import (
	"fmt"
	"os"

	c "common"
)

func Run(config c.YamlCode) {
	fmt.Println("Run python")
	file, err := os.Create("out.py")
	c.ChkErr(err)
	defer file.Close()
	for class, c := range config.YamlContent.(map[interface{}]interface{}) {
		WriteClass(file, class, c)
		fmt.Println(class)
		for f, v := range c.(map[interface{}]interface{}) {
			fmt.Println("\t", f)
			for l, d := range v.(map[interface{}]interface{}) {
				fmt.Println("\t\t", l)
				for g, s:= range d.(map[interface{}]interface{}) {
					fmt.Println("\t\t\t", g, ": ", s)
				}
			}
		}
	}
}

func WriteClass(file *os.File, class, c interface{}) {
	cl := fmt.Sprintf("class %s:\n", class)
	file.WriteString(cl)
	fmt.Println(cl)
	for f, v := range c.(map[interface{}]interface{}) {
		WriteFunc(file, f, v)
	}
}

func WriteFunc(file *os.File, fnc, v interface{}) {
	fn := fmt.Sprintf("\tdef %s(self):\n", fnc)
	file.WriteString(fn)
	fmt.Println(fn)
	for cb, d := range v.(map[interface{}]interface{}) {
		if (cb == "vars") {
		}
		switch cb {
		case "vars":
			for k, v := range d.(map[interface{}]interface{}) {
				WriteVar(file, k, v)
			}
		case "code":
			for k, v := range d.(map[interface{}]interface{}) {
				WriteCommands(file, k, v)
			}
		}
		fmt.Println(cb)
	}
}
	
func WriteVar(file *os.File, key, variable interface{}) {
	block := fmt.Sprintf("\t\t%s = \"%s\"\n", key, variable)
	file.WriteString(block)
	fmt.Println(block)
}

func WriteCommands(file *os.File, cmd, arg interface{}) {
	block := fmt.Sprintf("\t\t%s(%s)", cmd, arg)
	file.WriteString(block)
}