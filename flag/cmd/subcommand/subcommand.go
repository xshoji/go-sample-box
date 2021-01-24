package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// << Execution sample >>
// -----
// $ go run cmd/subcommand/subcommand.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build174983893/b001/exe/subcommand [commit|checkout]
//  -h
//    	help
//
//
//
// -----
// $ go run cmd/subcommand/subcommand.go commit -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build049923371/b001/exe/subcommand commit [options]
//  -a string
//    	[optional] commit all changed files
//  -h
//    	help
//  -m string
//    	[required] commit message
//
// $ go run cmd/subcommand/subcommand.go commit -m mmm -a aaa
// message:  mmm
// all:  aaa
//
//
//
// -----
// $ go run cmd/subcommand/subcommand.go checkout -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build206366477/b001/exe/subcommand checkout [options]
//  -b string
//    	[required] create and checkout a new branch
//  -h
//    	help
//  -p	[optional] force progress reporting
//
// $ go run cmd/subcommand/subcommand.go checkout -b bbb -p
// branch:  bbb
// progress:  true
//
func main() {
	flagSetMain := flag.NewFlagSet("main", flag.ContinueOnError)
	argsHelp := flagSetMain.Bool("h", false, "\nhelp")
	flagSetMain.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s [commit|checkout]\n", os.Args[0])
		flagSetMain.PrintDefaults()
	}
	if err := flagSetMain.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
	if *argsHelp || len(os.Args) < 2 {
		flagSetMain.Usage()
		os.Exit(0)
	}

	var err error
	switch os.Args[1] {
	case "commit":
		err = commit()
	case "checkout":
		err = checkout()
	default:
		err = fmt.Errorf("error: unknown command - %s", os.Args[1])
	}
	if err != nil {
		flagSetMain.Usage()
		log.Fatal(err)
	}

}

func commit() error {
	flagSetCommit := flag.NewFlagSet("commit", flag.ContinueOnError)
	var (
		argsMessage     = flagSetCommit.String("m", "", "[required] commit message")
		argsAll         = flagSetCommit.String("a", "", "[optional] commit all changed files")
		argsFlagSetHelp = flagSetCommit.Bool("h", false, "\nhelp")
	)
	flagSetCommit.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s commit [options]\n", os.Args[0])
		flagSetCommit.PrintDefaults()
	}
	if err := flagSetCommit.Parse(os.Args[2:]); err != nil {
		return err
	}
	if *argsFlagSetHelp || *argsMessage == "" {
		flagSetCommit.Usage()
		os.Exit(0)
	}

	fmt.Println("message: ", *argsMessage)
	fmt.Println("all: ", *argsAll)
	return nil
}

func checkout() error {
	flagSetCheckout := flag.NewFlagSet("checkout", flag.ContinueOnError)
	var (
		argsBranch      = flagSetCheckout.String("b", "", "[required] create and checkout a new branch")
		argsProgress    = flagSetCheckout.Bool("p", false, "[optional] force progress reporting")
		argsFlagSetHelp = flagSetCheckout.Bool("h", false, "\nhelp")
	)
	flagSetCheckout.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s checkout [options]\n", os.Args[0])
		flagSetCheckout.PrintDefaults()
	}
	if err := flagSetCheckout.Parse(os.Args[2:]); err != nil {
		return err
	}
	if *argsFlagSetHelp || *argsBranch == "" {
		flagSetCheckout.Usage()
		os.Exit(0)
	}

	fmt.Println("branch: ", *argsBranch)
	fmt.Println("progress: ", *argsProgress)
	return nil
}
