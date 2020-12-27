package main

import "fmt"

func main() {
	var array1 [5]int
	array1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("array1 = ", array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array2 = ", array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("array3 = ", array3)

}
