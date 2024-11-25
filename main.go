package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) > 1 {
		scriptFile := os.Args[1]
		script, err :=os.ReadFile(scriptFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading script")
			os.Exit(1)
		}
		repl.RunScript(string(script), os.Stdout)
		return
	}
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
