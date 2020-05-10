[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/schizoid90/code-as-config) 

# code-as-config
Is it possible to create working code from YAML?

This projects shows that it is possible and I'll worry about the useful later.

## Structure

YAML is used to define the code with a set structure. This structure is shown in `common/structs` but the below gives a good outline.

```yaml
version: v1 # version to work against (only v1 exists)
language: python # language to build (only python supported)
name: MyFirstScript # name of the resulting script
spec: # define the code
  class: MyClass # name of the class the code will live under
  runMain: "" # name of the function to call when running as main (i.e. from cli)
  functions: # define functions
  - name: myFunc # name of function
    vars: # list of variables to set up
      var1: "string" # key: value pair for the variables
    code: # commands to run, currently only supported in command(options) format
      command: option1, option2
    options: # add arguments to function (e.g. myFunc(opt="default")) !!Not currently supported!!
      opt: default # key: value pair for arguments, value is the default !!Not currently supported!!
```

See examples/simple_python.yaml for a working example. This will build a script that will output `hello world` when run with `python3 hello_world.py`

## TODO

* Allow for variable substitution
* imports
* More complicated scripts
    * custom code
    * loops
    * if/else/else if
* Multi file scripts

## Languages

* Python (3.7+)
