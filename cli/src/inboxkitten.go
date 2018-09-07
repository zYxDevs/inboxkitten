package main

import (
	"flag"
	"fmt"
	"os"
)

//
// Most the CLI example is taken from
//
// https://gobyexample.com/command-line-arguments
// https://gobyexample.com/command-line-flags
// https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/
//
func main() {

	// The api url with default value
	var apiDefault = "https://api.inboxkitten.com/api/v1/"
	var api string
	flag.StringVar(&api, "api", apiDefault, "URL to inbox kitten API")

	// Parse all the flags
	flag.Parse();
	
	// Output the api URL if its custom
	if( api != apiDefault ) {
		fmt.Printf("Using Custom API: %s \n", api)
	}

	// Post flag processing args
	var flagArgs = flag.Args();

	// Verify that a subcommand has been provided
	// flagArgs[0] is the main command
	// flagArgs[1] will be the subcommand
	var missingCommandError = "`list [email]` or `get [emailid]` subcommand is required\n";
	if len(flagArgs) <= 1 {
		fmt.Fprintf(os.Stderr, missingCommandError);
		flag.PrintDefaults();
		os.Exit(1);
	}

	// The list and get command
	getCommand := flag.NewFlagSet("get", flag.ExitOnError)
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// flagArgs[2:] will be all arguments starting after the subcommand at flagArgs[1]
	switch flagArgs[0] {
		case "list":
			listCommand.Parse(flagArgs[1:])
		case "get":
			getCommand.Parse(flagArgs[1:])
		default:
			fmt.Fprintf(os.Stderr, missingCommandError);
			flag.PrintDefaults();
			os.Exit(1);
	}


	// if len(image) == 0 {
	// 	fmt.Fprintf(os.Stderr, "You must specify a Docker image name")
	// }
	// fmt.Printf("Your Docker image was: %s", image)
	// fmt.Printf("\n")
}