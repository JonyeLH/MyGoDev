package main

import (
	"fmt"
	"strings"
)

func Index(data []string, strs string) int {
	for i, v := range data {
		if v == strs {
			return i
		}
	}
	return -1
}

func Include(data []string, strs string) bool {
	//for _,v := range data{
	//	if v == strs{
	//		return true
	//	}
	//}
	//return false
	//
	//利用Index函数来判断
	return Index(data, strs) > 0
}

func Any(data []string, f func(string) bool) bool {
	for _, v := range data {
		if f(v) {
			return true
		}
	}
	return false
}

func All(data []string, f func(string) bool) bool {
	for _, v := range data {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(data []string, f func(string) bool) []string {
	vaf := make([]string, 0)
	for _, v := range data {
		if f(v) {
			vaf = append(vaf, v)
		}
	}
	return vaf
}

func Map(data []string, f func(string) string) []string {
	vam := make([]string, len(data))
	for i, v := range data {
		vam[i] = f(v)
	}
	return vam
}

func main() {
	data := []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(data, "pear"))
	fmt.Println(Include(data, "grape"))
	fmt.Println(Any(data, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(data, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(data, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(data, strings.ToUpper)) //小写转大写
}
