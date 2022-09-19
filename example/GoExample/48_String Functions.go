package main

import "fmt"
import ss "strings" //strings取别名

var p = fmt.Println //给常用的 fmt.Println 取别名

func main() {

	p("Contains:  ", ss.Contains("Unit_Test", "es"))
	p("Count:     ", ss.Count("Unit_Test", "t"))
	p("HasPrefix: ", ss.HasPrefix("Unit_Test", "te"))
	p("HasSuffix: ", ss.HasSuffix("Unit_Test", "st"))
	p("Index:     ", ss.Index("Unit_Test", "e"))
	p("Join:      ", ss.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", ss.Repeat("a", 5))
	p("Replace:   ", ss.Replace("foo", "o", "0", -1))
	p("Replace:   ", ss.Replace("foo", "o", "0", 1))
	p("Split:     ", ss.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", ss.ToLower("TEST"))
	p("ToUpper:   ", ss.ToUpper("Unit_Test"))
}
