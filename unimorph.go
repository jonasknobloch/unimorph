package unimorph

import (
	"encoding/csv"
	"github.com/jonasknobloch/jinn/pkg/tree"
	"io"
	"os"
	"sort"
	"strconv"
	"time"
)

type Unimorph struct {
	lang string

	inflect map[string]map[string]struct{} // lemma -> form
	analyze map[string]map[string]struct{} // form -> lemma

	features map[string]map[string][]string // lemma -> form -> features
}

func (u *Unimorph) Init(dict string) error {
	f, err := os.Open(dict)

	if err != nil {
		return err
	}

	defer f.Close()

	r := csv.NewReader(f)

	r.Comma = '\t'

	u.inflect = make(map[string]map[string]struct{})
	u.analyze = make(map[string]map[string]struct{})

	u.features = make(map[string]map[string][]string)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		lemma, form, features := record[0], record[1], record[2]

		// initialize lookup tables

		if _, ok := u.inflect[lemma]; !ok {
			u.inflect[lemma] = make(map[string]struct{})
		}

		if _, ok := u.analyze[form]; !ok {
			u.analyze[form] = make(map[string]struct{})
		}

		if _, ok := u.features[lemma]; !ok {
			u.features[lemma] = make(map[string][]string)
		}

		// insert inflection data

		if _, ok := u.inflect[lemma][form]; !ok {
			u.inflect[lemma][form] = struct{}{}
		}

		if _, ok := u.analyze[form][lemma]; !ok {
			u.analyze[form][lemma] = struct{}{}
		}

		if _, ok := u.features[lemma][form]; !ok {
			u.features[lemma][form] = make([]string, 0, 1)
		}

		u.features[lemma][form] = append(u.features[lemma][form], features)
	}

	return nil
}

func (u *Unimorph) Inflect(lemma, features string) ([]string, bool) {
	if features != "" {
		return nil, false
	}

	forms, ok := u.inflect[lemma]

	if !ok {
		return nil, false
	}

	result := make([]string, 0)

	for form := range forms {
		if features == "" {
			result = append(result, form)

			continue
		}

		fss, _ := u.features[lemma][form]

		for _, fs := range fss {
			if fs == features {
				result = append(result, form)
			}
		}
	}

	return result, ok
}

func (u *Unimorph) Analyze(form string) ([]string, bool) {
	lemmas, ok := u.analyze[form]

	result := make([]string, 0, len(lemmas))

	for lemma := range lemmas {
		result = append(result, lemma)
	}

	return result, ok
}

func (u *Unimorph) Features(lemma, form string) ([]string, bool) {
	features, ok := u.features[lemma][form]

	return features, ok
}

func (u *Unimorph) Split(form string, f func(radix *tree.Tree) bool) ([]string, bool) {
	lemmas, ok := u.Analyze(form)

	if !ok {
		return nil, false
	}

	forms, ok := u.Inflect(lemmas[0], "")

	if !ok {
		return nil, false
	}

	if Config.SortRadixInput {
		sort.Strings(forms)
	}

	r := NewRadix()

	for _, form := range forms {
		r.Insert(form)
	}

	defer func() {
		if Config.CompressTrees {
			r.Compress()
		}

		if Config.DrawTrees {
			r.Draw(strconv.FormatInt(time.Now().Unix(), 10))
		}
	}()

	return r.Split(form, f), true
}
