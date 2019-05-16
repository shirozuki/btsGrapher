package main

import (
	"math"
	"math/rand"
	"time"
)

func statGen(radius float64, numOfStations int, stationRange float64) []btsPos {

	var x, y float64
	var listOfBts []btsPos

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < numOfStations; {

		x = (r1.Float64() * radius) - radius/2
		y = (r1.Float64() * radius) - radius/2

		if (math.Pow(x, 2) + math.Pow(y, 2)) < math.Pow(radius, 2) {
			newBts := btsPos{x, y}
			listOfBts = append(listOfBts, newBts)
			i++

		} else {
			continue
		}
	}

	return listOfBts
}
