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

	hits := float64(0)
	misses := float64(0)

	err := unimorph.Server("/tmp/unimorph.sock", func(form string) []byte {
		split, ok := u.Split(form, unimorph.SplitSiblings(2))

		if !ok {
			misses++

			fmt.Printf("%v\t%f\t%s\n", ok, hits/(hits+misses), form)

			return []byte(form)
		}

		hits++

		fmt.Printf("%v\t%f\t%s\n", ok, hits/(hits+misses), strings.Join(split, "#"))

		return []byte(strings.Join(split, "#"))
	})

	if err != nil {
		log.Fatalln(err)
	}
}
