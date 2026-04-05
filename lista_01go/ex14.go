package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a float64
	fmt.Scan(&h, &a)
	areaBase := (3 * a * a * math.Sqrt(3)) / 2
	volume := (areaBase * h) / 3
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", volume)
}
