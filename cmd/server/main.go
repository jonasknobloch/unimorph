package main

import (
	"fmt"
	"github.com/jonasknobloch/unimorph"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Starting server...")

	u := &unimorph.Unimorph{}

	fmt.Println("Initializing unimorph dictionary", "$HOME/.unimorph/ces/ces")

	if err := u.Init(filepath.Join(os.Getenv("HOME"), ".unimorph/ces/ces")); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Listing on socket", "/tmp/unimorph.sock")

	err := unimorph.Server("/tmp/unimorph.sock", func(form string) []byte {
		split, ok := u.Split(form, unimorph.SplitSiblings(2))

		if !ok {
			fmt.Printf("%v\t%s\n", ok, form)

			return []byte(form)
		}

		fmt.Printf("%v\t%s\n", ok, strings.Join(split, "#"))

		return []byte(strings.Join(split, "#"))
	})

	if err != nil {
		log.Fatalln(err)
	}
}
