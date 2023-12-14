package main

import (
	"github.com/jonasknobloch/unimorph"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	u := &unimorph.Unimorph{}

	if err := u.Init(filepath.Join(os.Getenv("HOME"), ".unimorph/ces/ces")); err != nil {
		log.Fatalln(err)
	}

	err := unimorph.Server("/tmp/unimorph.sock", func(form string) []byte {
		split, ok := u.Split(form, unimorph.SplitSiblings(2))

		if !ok {
			return []byte(form)
		}

		return []byte(strings.Join(split, " "))
	})

	if err != nil {
		log.Fatalln(err)
	}
}
