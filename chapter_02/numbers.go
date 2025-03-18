package chapter_02

import "fmt"

func Numbers() {
	// Complex numbers are directly supported by go
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)
	var c3 = complex64(c1 + c2)
	fmt.Println("c3:", c3)
	fmt.Printf("Type of c3: %T\n", c3)
	cZero := c3 - c3
	fmt.Println("cZero:", cZero)

	// As per usual: int / int = int. Keep in mind when letting go decide the type of the variable.
	// int = int64 or int32 depending on cpu
	x := 12
	k := 5
	fmt.Println("x:", x)
	fmt.Printf("Type of x: %T\n", x)
	div := x / k
	fmt.Println("div 12 / 5 =", div)

	// Can be avoided by explicitly declaring the type
	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n)
	y := 4 / 2.3
	fmt.Println("y:", y)
	// Or by converting the integers
	divFloat := float64(x) / float64(k)
	fmt.Println("divFloat 12 / 5:", divFloat)
	fmt.Printf("Type of divFloat: %T\n", divFloat)
}
