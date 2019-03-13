package main

import (
	"fmt"
	"math"
)

func conflictDetection(listOfBts []btsPos, stationRange float64) {

	for i := 0; i < len(listOfBts); i++ {
		for j := i + 1; j < len(listOfBts); j++ {
			if i == j {
				continue
			} else {
				distance := math.Sqrt(math.Pow(listOfBts[j].x-listOfBts[i].x, 2) + math.Pow(listOfBts[j].y-listOfBts[i].y, 2))

				if distance <= stationRange*2 {
					fmt.Printf("Stacja %d oraz %d - TAK\n", i+1, j+1)
				} else {
					fmt.Printf("Stacja %d oraz %d - NIE\n", i+1, j+1)
				}
			}
		}
	}
}
