package main

import "fmt"

func main() {
	//first name
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	//second name
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	delete(m, "k2")
	fmt.Println("map:", m)

	m["k3"] = 3
	fmt.Println("map:", m)

	for _, sec := range m {
		fmt.Println("sec:", sec)
	}

	//map type slice
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
}
