package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// ここをカスタムする
func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return l[i].Name < l[j].Name
	} else {
		return l[i].Value < l[j].Value
	}
}

// sort
func main() {
	/*
		i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10}
		s := []string{"a", "z", "j"}

		fmt.Println(i, s)
		sort.Ints(i)
		sort.Strings(s)
		fmt.Println(i, s)
	*/

	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}

	fmt.Println(el)

	// Slice
	/*
		sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })
		sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })
		fmt.Println(el)
	*/

	// SliceStable
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	fmt.Println(el)

	m := map[string]int{"S": 1, "J": 4, "A": 3, "N": 3}
	lt := List{}

	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

	// Sort
	sort.Sort(lt)

	fmt.Println(lt)

	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)
}
