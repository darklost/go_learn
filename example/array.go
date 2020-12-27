//定义数组
package main

import "fmt"

func main() {
	//初始化数组
	var array1 [5]int
	array1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("array1 = ", array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array2 = ", array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("array3 = ", array3)

	array4 := [5]int{0, 1, 0, 4, 0}
	fmt.Println("array4 = ", array4)

	array5 := [5]int{1: 1, 3: 4}
	fmt.Println("array4 = ", array5)

	//修改数组
	array6 := [5]int{1: 1, 3: 4}
	fmt.Printf("%d\n", array6[1])
	array6[1] = 3
	fmt.Printf("%d\n", array6[1])

	//遍历数组
	for i := 0; i < 5; i++ {
		fmt.Printf("索引:%d,值:%d\n", i, array6[i])
	}

	for i, v := range array6 {
		fmt.Printf("索引:%d,值:%d\n", i, v)
	}

	//多维数组

	var tdArray = [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	fmt.Println("tdArray = ", tdArray)
	var tdArray1 = [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}} /* 第三行索引为 2 */
	fmt.Println("tdArray1 = ", tdArray1)
	//访问二维数组
	val := tdArray[2][3]
	fmt.Println("val = ", val)
	// 或
	var value int = tdArray[2][3]
	fmt.Println("value = ", value)
}
