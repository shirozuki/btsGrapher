package main

import (
	"fmt"
	"os"
	"os/exec"
)

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
		fmt.Println("Stacja", i+1, ":", listOfBts[i])
	}

	conflictMap := conflictDetection(listOfBts, stationRange)

	//fmt.Println(conflictMap)

	generateFile(numOfStations, listOfBts, conflictMap)

	fmt.Println("-------------------------------------------------")

	cmd := exec.Command("python", "./assignFreq.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	os.Remove("graph.connections")
	os.Remove("graph.plain")

}
