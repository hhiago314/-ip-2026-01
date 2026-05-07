package main

import "fmt"

func main() {
	var v [100]int
	for i := 0; i < 100; i++ {
		v[i] = 2*i + 1
	}
	fmt.Println(v)
}
