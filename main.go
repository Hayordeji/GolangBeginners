package main

import "fmt"

func main() {
	//ARRAY INITIALIZATION
	var array1 = [5]int{1, 2, 3, 4, 0}

	size := 4
	firstNumber := 10

	//INSERT ELEMENT AT THE BEGINNING OF THE ARRAY
	for i := size; i > 0; i-- {
		array1[i] = array1[i-1]
	}
	array1[0] = firstNumber

	fmt.Println("New array after inserting is : ", array1)

	//INSERTING AT THE LAST ELEMENT
	array2 := [6]int{0, 2, 3, 4, 5, 6}
	size = (len(array2) - 1)
	lastNumber := 22
	for i := 0; i < size; i++ {
		array2[i] = array2[i+1]
	}

	array2[len(array2)-1] = lastNumber
	fmt.Println("Array 2 after inserting at the last eleement is : ", array2)

	//INSERTING AT A PARTICULAR INDEX
	array3 := [6]int{0, 2, 3, 4, 5, 6}
	elementIndex := 2
	numberToAdd := 100

	for i := 0; i < elementIndex; i++ {
		array3[i] = array3[i+1]
	}
	array3[elementIndex] = numberToAdd
	fmt.Println("Array 3 after inserting at the last eleement is : ", array3)
}
