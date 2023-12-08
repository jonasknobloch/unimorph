package unimorph

import (
	"fmt"
	"github.com/jonasknobloch/jinn/pkg/tree"
	"os"
	"path/filepath"
	"strings"
)

type Radix tree.Tree

func NewRadix() *Radix {
	return (*Radix)(&tree.Tree{
		Label:    "ROOT", // empty word
		Children: []*tree.Tree{},
	})
}

func (r *Radix) Insert(s string) []int {
	rs := []rune(s + "$")

	return insert((*tree.Tree)(r), rs, make([]int, len(rs)))
}

func (r *Radix) Split(s string, f func(t *tree.Tree) bool) []string {
	rs := []rune(s + "$")

	result := make([]string, 0)

	for _, rs := range split((*tree.Tree)(r), rs, make([]rune, 0), make([][]rune, 0), f) {
		result = append(result, string(rs))
	}

	return result
}

func insert(t *tree.Tree, rs []rune, p []int) []int {
	if len(rs) == 0 {
		return p
	}

	for i, c := range t.Children {
		if c.Label == string(rs[0]) {
			return insert(c, rs[1:], append(p, i))
		}
	}

	if t.Children == nil {
		t.Children = make([]*tree.Tree, 0, 1)
	}

	t.Children = append(t.Children, &tree.Tree{Label: string(rs[0])})

	return insert(t, rs, p)
}

func split(t *tree.Tree, rs []rune, s []rune, r [][]rune, f func(t *tree.Tree) bool) [][]rune {
	if len(rs) == 1 {
		if rs[0] != '$' {
			panic("invalid end of word")
		}

		ok := false

		for _, c := range t.Children {
			if c.Label == "$" {
				ok = true
				break
			}
		}

		if !ok {
			panic("unexpected end of word")
		}

		if len(s) != 0 {
			r = append(r, s)
		}

		return r
	}

	if len(rs) > 1 && len(t.Children) == 0 {
		panic("unexpected end of path")
	}

	for _, c := range t.Children {
		if c.Label == string(rs[0]) {
			s = append(s, rs[0])

			if f(c) {
				r = append(r, s)    // add rune to substring
				s = make([]rune, 0) // reset substring
			}

			return split(c, rs[1:], s, r, f)
		}
	}

	return split(t, rs, s, r, f)
}

func (r *Radix) Compress() {
	t := (*tree.Tree)(r)

	if len(t.Children) == 0 {
		return
	}

	compress(t.Children[0])
}

func compress(t *tree.Tree) {
	if len(t.Children) == 0 {
		return
	}

	if len(t.Children) == 1 {
		if t.Children[0].Label == "$" {
			return
		}

		t.Label = string(append([]rune(t.Label), []rune(t.Children[0].Label)[0]))
		t.Children = t.Children[0].Children

		compress(t)
	}

	for _, c := range t.Children {
		compress(c)
	}
}

func (r *Radix) Draw(stubs ...string) (int, error) {
	t := (*tree.Tree)(r)

	name := fmt.Sprintf("radix_%s.dot", strings.Join(stubs, "-"))

	f, err := os.Create(filepath.Join(name))

	if err != nil {
		return 0, err
	}

	defer f.Close()

	sb := strings.Builder{}

	sb.WriteString("digraph D {\n")
	sb.WriteString("  node [shape=record]\n")

	t.Walk(func(st *tree.Tree) {
		sb.WriteString(fmt.Sprintf("  PTR%p", st))
		sb.WriteString(" [")

		if len(st.Children) > 0 {
			sb.WriteString("color=blue label=\"")
		} else {
			sb.WriteString("color=red label=\"")
		}

		sb.WriteString(fmt.Sprintf("%s ", st.Label))

		sb.WriteString("\"]\n")
	})

	for k, cs := range t.Edges() {
		for _, c := range cs {
			sb.WriteString(fmt.Sprintf("  PTR%p -> PTR%p\n", k, c))
		}
	}

	sb.WriteString("}\n")

	return f.WriteString(sb.String())
}
