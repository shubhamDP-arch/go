package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jaqen/goober/src/commands"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Error you fucker add the stupid command like this goober <command> -flags")
	}

	command := os.Args[1]

	switch command {
	case "init":
		commands.Init()
	case "add":
		commands.Add()
	case "commit":
		//this is horseshit
		fooCmd := flag.NewFlagSet("commit", flag.ExitOnError)
		message := fooCmd.String("m", "foo", "Commit message")
		fooCmd.Parse(os.Args[2:])

		if *message == "" {
			log.Fatal("Error idiot ")
		}
		fmt.Printf("Message: %s\n", *message)

		commands.Commit(*message)
	}

}
