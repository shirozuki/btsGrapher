package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func statGen() []btsPos {

	var radius, stationRange float64
	var numOfStations int
	fmt.Printf("Promień: ")
	fmt.Scan(&radius)
	fmt.Printf("Liczba stacji: ")
	fmt.Scan(&numOfStations)
	fmt.Printf("Zasięg: ")
	fmt.Scan(&stationRange)

	// Środek okręgu ustalony na [0,0]

	// koncept: random generujący się w zależności od podanego R.
	// 			niech generuje się w kwadracie, czyli na osi współrzędnych.
	//			Następnie sprawdzenie, czy punkt nie padł za daleko od środka koła (czy odległość między punktem a środkiem nie jest większa niż r)

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
