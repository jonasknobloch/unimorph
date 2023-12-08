package unimorph

import (
	"fmt"
	"os"
	"sort"
	"testing"
)

var unimorph *Unimorph

var lemma string
var form string

func TestMain(m *testing.M) {
	unimorph = &Unimorph{}

	err := unimorph.Init("test/ces_afghansky.tsv")

	if err != nil {
		return
	}

	lemma = "afghánský"
	form = "afghánskou"

	code := m.Run()
	os.Exit(code)
}

func ExampleUnimorph_Inflect() {
	forms, ok := unimorph.Inflect(lemma, "")

	if !ok {
		return
	}

	sort.Strings(forms)

	fmt.Println(forms)
	// Output: [afghánskou afghánská afghánské afghánského afghánském afghánskému afghánský afghánských afghánským afghánskými afghánští]
}

func ExampleUnimorph_Analyze() {
	lemmas, ok := unimorph.Analyze(form)

	if !ok {
		return
	}

	sort.Strings(lemmas)

	fmt.Println(lemmas)
	// Output: [afghánský]
}

func ExampleUnimorph_Features() {
	features, ok := unimorph.Features(lemma, form)

	if !ok {
		return
	}

	fmt.Println(features)
	// Output: [ADJ;ACC;FEM;SG]
}

func ExampleUnimorph_Split() {
	split, ok := unimorph.Split(form, SplitSiblings(3))

	if !ok {
		return
	}

	fmt.Println(split)
	// Output: [afghánsk ou]
}
