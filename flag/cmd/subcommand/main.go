package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// << Execution sample >>
// -----
// $ go run cmd/subcommand/main.go -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build174983893/b001/exe/subcommand [commit|checkout]
//
//	-h
//	  	help
//
// -----
// $ go run cmd/subcommand/main.go commit -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build049923371/b001/exe/subcommand commit [options]
//
//	-a string
//	  	[optional] commit all changed files
//	-h
//	  	help
//	-m string
//	  	[required] commit message
//
// $ go run cmd/subcommand/main.go commit -m mmm -a aaa
// message:  mmm
// all:  aaa
//
// -----
// $ go run cmd/subcommand/main.go checkout -h
// Usage of /var/folders/_q/dpw924t12bj25568xfxcd2wm0000gn/T/go-build206366477/b001/exe/subcommand checkout [options]
//
//	-b string
//	  	[required] create and checkout a new branch
//	-h
//	  	help
//	-p	[optional] force progress reporting
//
// $ go run cmd/subcommand/main.go checkout -b bbb -p
// branch:  bbb
// progress:  true
func main() {
	flagSetMain := flag.NewFlagSet("main", flag.ContinueOnError)
	paramsHelp := flagSetMain.Bool("h", false, "\nhelp")
	flagSetMain.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s [commit|checkout]\n", os.Args[0])
		flagSetMain.PrintDefaults()
	}
	if err := flagSetMain.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
	if *paramsHelp || len(os.Args) < 2 {
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
		paramsMessage     = flagSetCommit.String("m", "", "[required] commit message")
		paramsAll         = flagSetCommit.String("a", "", "[optional] commit all changed files")
		paramsFlagSetHelp = flagSetCommit.Bool("h", false, "\nhelp")
	)
	flagSetCommit.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s commit [options]\n", os.Args[0])
		flagSetCommit.PrintDefaults()
	}
	if err := flagSetCommit.Parse(os.Args[2:]); err != nil {
		return err
	}
	if *paramsFlagSetHelp || *paramsMessage == "" {
		flagSetCommit.Usage()
		os.Exit(0)
	}

	fmt.Println("message: ", *paramsMessage)
	fmt.Println("all: ", *paramsAll)
	return nil
}

func checkout() error {
	flagSetCheckout := flag.NewFlagSet("checkout", flag.ContinueOnError)
	var (
		paramsBranch      = flagSetCheckout.String("b", "", "[required] create and checkout a new branch")
		paramsProgress    = flagSetCheckout.Bool("p", false, "[optional] force progress reporting")
		paramsFlagSetHelp = flagSetCheckout.Bool("h", false, "\nhelp")
	)
	flagSetCheckout.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s checkout [options]\n", os.Args[0])
		flagSetCheckout.PrintDefaults()
	}
	if err := flagSetCheckout.Parse(os.Args[2:]); err != nil {
		return err
	}
	if *paramsFlagSetHelp || *paramsBranch == "" {
		flagSetCheckout.Usage()
		os.Exit(0)
	}

	fmt.Println("branch: ", *paramsBranch)
	fmt.Println("progress: ", *paramsProgress)
	return nil
}
