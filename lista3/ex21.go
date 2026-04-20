package main

import "fmt"

func main() {
    var b, n int
    fmt.Print("Digite b (>=2) e n (>1): ")
    fmt.Scan(&b, &n)

    resultado := 1
    for i := 0; i < n; i++ {
        resultado *= b
    }
    fmt.Printf("%d^%d =