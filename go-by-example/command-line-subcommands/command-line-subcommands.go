package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Call with e.g.
	// ./command-line-subcommands/build/command-line-subcommands cmd1 -enable --name=joe a1 a2
	// ./command-line-subcommands/build/command-line-subcommands cmd2 --level=15 a4

	fooCmd := flag.NewFlagSet("cmd1", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	barCmd := flag.NewFlagSet("cmd2", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'cmd1' or 'cmd2' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "cmd1":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'cmd1'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "cmd2":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'cmd2'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'cmd1' or 'cmd2' subcommands")
		os.Exit(1)
	}
}
