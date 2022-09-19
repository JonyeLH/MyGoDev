package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"a", "c", "b"}
	sort.Strings(strs)
	fmt.Println("strs is sort:", strs)
	strsortted := sort.StringsAreSorted(strs)
	fmt.Println("strs is sortted:", strsortted)

	ints := []int{1, 3, 2}
	sort.Ints(ints)
	fmt.Println("ints is sort:", ints)
	Intsortted := sort.IntsAreSorted(ints)
	fmt.Println("ints is sortted:", Intsortted)
}
