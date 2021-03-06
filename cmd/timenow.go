package main

import (
	"fmt"
	"os"

	"github.com/jodosha/timenow"
)

func main() {
	t := timenow.New(nil)
	// t := timenow.New(http.DefaultClient)

	now, err := t.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(now)
}
