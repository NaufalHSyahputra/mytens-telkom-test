package main

import (
	"fmt"
	"os"
	"path"

	"github.com/thatisuday/commando"
)

func main() {

	// set CLI executable, version and description
	commando.
		SetExecutableName("mytest-test").
		SetVersion("v1.0.0").
		SetDescription("Tools ini berfungsi untuk mengambil file log pada system milik Linux pada folder /var/log")

	// configure the root-command
	// $ mytenst-test <filename>  --type|-t  --output|-o --version|-v  --help|-h
	commando.
		Register(nil).
		AddArgument("filename", "Log File Location", "").                                 // required
		AddFlag("type, t", "Convert log to plaintext or JSON ", commando.String, "text"). // optional
		AddFlag("output, o", "Set new output location", commando.String, "same").         // optional
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fileLocation := args["filename"].Value
			newFileLocation := flags["output"].Value
			newFileLocationString := fmt.Sprintf("%v", newFileLocation)
			typeFlag := flags["type"].Value
			typeString := fmt.Sprintf("%v", typeFlag)
			if typeString == "text" {
				typeString = "txt"
			}
			if newFileLocation != "same" {
				ext := path.Ext(newFileLocationString)
				outfile := newFileLocationString[0:len(newFileLocationString)-len(ext)] + "." + typeString
				_, isError := doCopy(fileLocation, outfile)
				if isError != nil {
					fmt.Println("Process Error: " + isError.Error())
					os.Exit(3)
				}
				fmt.Println("Process Success!")
			} else {
				ext := path.Ext(fileLocation)
				outfile := fileLocation[0:len(fileLocation)-len(ext)] + "." + typeString
				_, isError := doCopy(fileLocation, outfile)
				if isError != nil {
					fmt.Println("Process Error: " + isError.Error())
					os.Exit(3)
				}
				fmt.Println("Process Success!")
			}
		})
	// parse command-line arguments from the STDIN
	commando.Parse(nil)
}
