package main

import (
	"bitbucket.org/instinctools/gluten/cli/client"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"flag"
	"os"
)

func main() {
	//lists command
	runOwnerCommand := flag.NewFlagSet("run", flag.ExitOnError)
	generateOwnerCommand := flag.NewFlagSet("generate", flag.ExitOnError)
	//flag for help
	helpOwnerCommand := flag.NewFlagSet("help", flag.ExitOnError)

	//run command
	masterCommand := runOwnerCommand.String("m", "", "-master, [address:port]")
	pathToFileCommand := runOwnerCommand.String("pf", "", "-path_to_file, [filename]."+
		" If the file path is not specified, it will use the standard path ")

	//generate command
	autoGenerateFileCommand := generateOwnerCommand.String("o", "", "-output_file [filename]")

	parseFlagsCommand(runOwnerCommand, generateOwnerCommand, helpOwnerCommand)

	if runOwnerCommand.Parsed() {
		if *masterCommand == "" || *pathToFileCommand == "" {
			runOwnerCommand.PrintDefaults()
			os.Exit(1)
		} else {
			json := ReadJSONFile(*pathToFileCommand)
			obj.LaunchClient(*masterCommand, json)
			os.Exit(1)
		}
	}

	if generateOwnerCommand.Parsed() {
		if *autoGenerateFileCommand == "" {
			generateOwnerCommand.PrintDefaults()
			os.Exit(1)
		} else {
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

func NilHandler(err error) {
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Can't connect to master")

		log.Fatal("Error", err)
	}
}

func parseFlagsCommand(runOwnerCommand *flag.FlagSet, generateOwnerCommand *flag.FlagSet, helpOwnerCommand *flag.FlagSet) {
	if len(os.Args) < 2 {
		fmt.Println("Command is wrong. Pls try again")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "run":
		err = runOwnerCommand.Parse(os.Args[2:])
		NilHandler(err)
	case "generate":
		err = generateOwnerCommand.Parse(os.Args[2:])
		NilHandler(err)
	case "help":
		err = helpOwnerCommand.Parse(os.Args[2:])
		NilHandler(err)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
