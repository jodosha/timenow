package main

import (
	"fmt"
	"os"

	"github.com/jodosha/timenow"
)

func main() {
	now, err := timenow.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(now)
}
