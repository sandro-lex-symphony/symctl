package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sandro-lex-symphony/symctl"
)

func usage() {
	fmt.Printf("Usage:\n" +
		"\tcreate-users [-n count] [-o file]\n")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	numberOfUsers := flag.Int("n", 1, "Number of users to generate")
	output := flag.String("o", "", "Output file")

	flag.Parse()

	switch flag.Arg(0) {
	case "create-users":
		usersArray := symctl.CreateUsers(*numberOfUsers)
		if *output != "" {
			symctl.WriteCSV(*output, usersArray)
		} else {
			symctl.ShowCSV(usersArray)

		}
	case "test":
		fnames := symctl.GetNames("data/1.txt")
		fmt.Println(fnames)

	}

}
