package main

// https://qiita.com/taku-yamamoto22/items/4d6f9ff8451a0b86997b

import (
	"fmt"
	f"fmt"
	."fmt"
	"golang_udemy/lesson54/foo"
)

// スコープ

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func main() {
	fmt.Println(foo.Max)
	// fmt.Println(foo.min)

	f.Println(foo.ReturnMin())
	Println(foo.Max)

	fmt.Println(appName())
	// fmt.Println(AppName, Version)

	fmt.Println(Do("AAA"))
}
