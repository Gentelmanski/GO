package main

import "fmt"

func main() {
	/*score := 111
	score = 50
	test := "game over"

	fmt.Println(test)
	fmt.Println(score)
	fmt.Println("Test")*/
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}

	numbers := []int{1, 2, 3, 4, 5}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

}
