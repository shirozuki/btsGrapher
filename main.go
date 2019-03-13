package main

import "fmt"

func main() {

	listOfBts := statGen()

	for i := 0; i < len(listOfBts); i++ {
		fmt.Println(listOfBts[i])
	}

	conflictDetection(listOfBts, 1)

}
