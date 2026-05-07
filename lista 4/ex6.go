package main

import "fmt"

func main() {
	var v [100]int
	for i := 0; i < 100; i++ {
		v[i] = 100 - i
	}
	fmt.Println(v)
}
