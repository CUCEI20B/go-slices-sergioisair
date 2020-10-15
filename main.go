package main

import "fmt"

func main() {
	var n int
	var x int
	var suma int
	fmt.Scan(&n)
	var s []int
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		s = append(s, x)
	}
	for i := 0; i < n; i++ {
		suma = suma + s[i]
	}
	fmt.Print(suma)
}
