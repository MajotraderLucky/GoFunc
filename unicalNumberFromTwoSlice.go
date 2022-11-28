// Принимаем два интовых числа, возвращаем два слайса, выводим
// числа, которые есть в обоих срезах.
package main

import "fmt"

func number2slise(a int) []int {
	var faSlice []int
	for i := a; i > 0; i /= 10 {
		b := i % 10
		faSlice = append(faSlice, b)
	}
	return faSlice
}

func contains(a []int, b int) []int {
	var unicNumbersSlice []int
	for _, n := range a {
		if b == n {
			unicNumbersSlice = append(unicNumbersSlice, b)
		}
	}
	return unicNumbersSlice
}

func main() {
	a := 5649
	b := 8954
	aSlice := number2slise(a)
	bSlice := number2slise(b)

	var unic []int

	fmt.Println(aSlice)
	fmt.Println(bSlice)

	for i := len(aSlice) - 1; i >= 0; i-- {
		unicNSlice := contains(bSlice, aSlice[i])
		unic = append(unic, unicNSlice...)
	}

	for i := 0; i < len(unic); i++ {
		if i != len(unic)-1 {
			fmt.Print(unic[i], " ")
		} else {
			fmt.Print(unic[i])
		}
	}
	fmt.Println("")
}
