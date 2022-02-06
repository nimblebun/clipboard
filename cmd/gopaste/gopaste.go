package main

import (
	"fmt"

	"pkg.nimblebun.works/clipboard"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Print(text)
}
