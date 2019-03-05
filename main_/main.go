package main

import (
	"fmt"
	"math"
)

func main() {

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(math.Sin(n))
}
