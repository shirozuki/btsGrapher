package main

import "fmt"

func main() {

	var radius, stationRange float64
	var numOfStations int
	fmt.Printf("Promień: ")
	fmt.Scan(&radius)
	fmt.Printf("Liczba stacji: ")
	fmt.Scan(&numOfStations)
	fmt.Printf("Zasięg: ")
	fmt.Scan(&stationRange)

	listOfBts := statGen(radius, numOfStations, stationRange)

	for i := 0; i < len(listOfBts); i++ {
		fmt.Println("Stacja", i+1, listOfBts[i])
	}

	conflictMap := conflictDetection(listOfBts, stationRange)

	fmt.Println(conflictMap)

	generateFile(numOfStations, listOfBts, conflictMap)

}
