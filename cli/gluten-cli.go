package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	obj "./client"
)

// Create a new type for a list of Strings
type stringList []string

// Implement the flag.Value interface
func (s *stringList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

func main() {
	//lists command
	runOwnerCommand := flag.NewFlagSet("run", flag.ExitOnError)
	generateOwnerCommand := flag.NewFlagSet("generate", flag.ExitOnError)
	//flag for help
	helpOwnerCommand := flag.NewFlagSet("help", flag.ExitOnError)

	//run command
	masterCommand := runOwnerCommand.String("m", "", "-master, [address:port]")
	pathToFileCommand := runOwnerCommand.String("pf", "", "-path_to_file, [filename]." +
		" If the file path is not specified, it will use the standard path ")

	//generate command
	autoGenerateFileCommand := generateOwnerCommand.String("o", "", "-output_file [filename]")

	if len(os.Args) < 2 {
		fmt.Println("Command is wrong. Try again")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runOwnerCommand.Parse(os.Args[2:])
	case "generate":
		generateOwnerCommand.Parse(os.Args[2:])
	case "help":
		helpOwnerCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if runOwnerCommand.Parsed() {
		if *masterCommand == "" || *pathToFileCommand == "" {
			runOwnerCommand.PrintDefaults()
			os.Exit(1)
		}else {
			json := ReadJsonFile(*pathToFileCommand)
			obj.LaunchClient(*masterCommand, json)
			os.Exit(1)
		}
	}

	if generateOwnerCommand.Parsed() {
		if *autoGenerateFileCommand == "" {
			generateOwnerCommand.PrintDefaults()
			os.Exit(1)
		}else {
			AutoGenerateConfig(*autoGenerateFileCommand)
			os.Exit(1)
		}
	}

	if helpOwnerCommand.Parsed() {
		println("For 'run' command:")
		runOwnerCommand.PrintDefaults()
		println("For 'generate' command:")
		generateOwnerCommand.PrintDefaults()
		os.Exit(1)
	}
}