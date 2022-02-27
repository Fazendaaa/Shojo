package shojo

import (
	"fmt"
	"os"
)

// Stores the description related to given command
type Usage struct {
	Short string
	Long  string
}

type CMD struct {
	Name     string
	Usage    Usage
	Params   []string // Just the parameters name
	Function func(...interface{})
}

var ShojoCMD = []CMD{
	{
		Name: "init",
		Usage: Usage{
			Short: "Initialize a new project",
			Long:  "",
		},
		Params: []string{
			"project path",
		},
		Function: func(params ...interface{}) {
		},
	},
	{
		Name: "add",
		Usage: Usage{
			Short: "Adds a LaTex package to the current project",
			Long:  "",
		},
		Params: []string{
			"package name",
		},
		Function: func(params ...interface{}) {
			// fmt.Println(params)
			fmt.Println("FOO")
			path, fail := os.Getwd()
			fmt.Println("BAR")
			// fmt.Println(params)

			fmt.Println(path)

			if nil != fail {
				fmt.Errorf("%w;\ncould not read current directory", fail)
			}

			project, fail := Load(path)

			fmt.Println(project)

			if nil != fail {
				fmt.Errorf("%w;\nmalformed tex definition", fail)
			}
		},
	},
	{
		Name: "rm",
		Usage: Usage{
			Short: "Removes a LaTex package to the current project",
			Long:  "",
		},
		Params: []string{
			"package name",
		},
		Function: func(params ...interface{}) {
			fmt.Println(params)
		},
	},
}
