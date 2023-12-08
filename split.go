package unimorph

import "github.com/jonasknobloch/jinn/pkg/tree"

func NoSplit() func(t *tree.Tree) bool {
	return func(_ *tree.Tree) bool {
		return false
	}
}
func SplitCharacters() func(t *tree.Tree) bool {
	return func(_ *tree.Tree) bool {
		return true
	}
}
func SplitSiblings(n int) func(t *tree.Tree) bool {
	return func(t *tree.Tree) bool {
		return len(t.Children) > n
	}
}
