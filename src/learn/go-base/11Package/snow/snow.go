package snow

import (
	"11Package/calc"
	"fmt"
)

func Test(x, y int) {
	fmt.Printf("x + y = %d\n", calc.Add(x, y))
	fmt.Printf("x - y = %d\n", calc.Sub(x, y))
	fmt.Printf("x * y = %d\n", calc.Mul(x, y))
	fmt.Printf("x / y = %d\n", calc.Div(x, y))
}
