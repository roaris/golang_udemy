package main

import "fmt"

// マップ

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3])

	_, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	// fmt.Println(s, ok)

	m4[2] = "US"
	fmt.Println(m4)

	m4[3] = "CHINA"
	fmt.Println(m4)

	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))
}
