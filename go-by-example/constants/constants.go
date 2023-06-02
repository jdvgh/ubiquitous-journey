package main

import (
	"fmt"
	"math"
)

const s string = "constant_string"
const i int = 1
const f float64 = 1.55

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(i)

	fmt.Println(f)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
