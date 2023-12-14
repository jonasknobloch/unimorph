package main

import (
	"fmt"
	"github.com/jonasknobloch/unimorph"
	"os"
	"path/filepath"
)

func main() {
	u := &unimorph.Unimorph{}

	if err := u.Init(filepath.Join(os.Getenv("HOME"), ".unimorph/ces/ces")); err != nil {
		return
	}

	form := "afghánský"

	if len(os.Args) > 2 {
		form = os.Args[1]
	}

	split, ok := u.Split(form, unimorph.SplitSiblings(2))

	if !ok {
		fmt.Println([]string{form})

		return
	}

	fmt.Println(split)
}
