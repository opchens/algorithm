package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	fmt.Println(m)
	fmt.Println(m)
	for k, v := range m {

		println(k, v)
	}

	for k, v := range m {
		println(k, v)
	}

}
