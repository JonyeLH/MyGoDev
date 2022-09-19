package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("abc")
	fmt.Println(string(strB))
	slcB := []string{"apple", "peach", "pear"}
	slcD, _ := json.Marshal(slcB)
	fmt.Println(string(slcD))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapC, _ := json.Marshal(mapD)
	fmt.Println(string(mapC))

	res1B := &response1{
		Page:   1,
		Fruits: []string{"apple", "pear"},
	}
	res1D, _ := json.Marshal(res1B)
	fmt.Println(string(res1D))
	res2B := &response2{
		Page:   1,
		Fruits: []string{"apple", "pear", "peach"},
	}
	res2D, _ := json.Marshal(res2B)
	fmt.Println(string(res2D))

	byt := []byte(`{"num": 6.13, "strs": ["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
