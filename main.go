package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey interpreter!\n", user.Username)
	fmt.Printf("Type in some commands:\n")
	repl.Start(os.Stdin)
}
