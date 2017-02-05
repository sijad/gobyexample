// `for` تنها دستور حلقه در Go است. اینجا
// سه حالت مختلف `for` آورده شده است.

package main

import "fmt"

func main() {

	// ساده‌ترین حالت، تنها با یک شرط.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// حالت مقداردهی‌اولیه-شرط-بروزرسانی حلقه `for`.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// حلقه `for` بدون شرط تا وقتی که دستور `break` و یا
	// `return` صدا شوند مکررا تکرار می‌شود.
	for {
		fmt.Println("loop")
		break
	}
}
