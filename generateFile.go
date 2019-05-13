package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generateFile(numOfStations int, listOfBts []btsPos, conflictMap []Conflict) {

	var buffer []string

	file, err := os.Create("graph.xml")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	buffer = append(buffer, "<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n")
	buffer = append(buffer, "<!DOCTYPE gxl SYSTEM \"http://www.gupro.de/GXL/gxl-1.0.dtd\">\n")
	buffer = append(buffer, "<gxl>\n")
	buffer = append(buffer, "<graph>\n")

	for i := 0; i < numOfStations; i++ {
		buffer = append(buffer, "<node id=\"")
		nodeID := strconv.Itoa(i + 1)
		buffer = append(buffer, nodeID)
		buffer = append(buffer, "\">\n")
		buffer = append(buffer, "<x>")
		nodeXPos := strconv.FormatFloat(listOfBts[i].x, 'f', 6, 64)
		buffer = append(buffer, nodeXPos)
		buffer = append(buffer, "</x>\n")
		buffer = append(buffer, "<y>")
		nodeYPos := strconv.FormatFloat(listOfBts[i].y, 'f', 6, 64)
		buffer = append(buffer, nodeYPos)
		buffer = append(buffer, "</y>\n")
		buffer = append(buffer, "</node>\n")
	}

	for i := 0; i < len(conflictMap); i++ {
		tmp := strconv.Itoa(i + 1)
		from := strconv.Itoa(conflictMap[i].a + 1)
		to := strconv.Itoa(conflictMap[i].b + 1)
		buffer = append(buffer, "<edge id=\"")
		buffer = append(buffer, tmp)
		buffer = append(buffer, "\"")
		buffer = append(buffer, " from=\"")
		buffer = append(buffer, from)
		buffer = append(buffer, "\"")
		buffer = append(buffer, " to=\"")
		buffer = append(buffer, to)
		buffer = append(buffer, "\"")
		buffer = append(buffer, " />\n")
	}

	buffer = append(buffer, "</graph>\n")
	buffer = append(buffer, "</gxl>\n")
	fmt.Println(buffer)

	outputStr := strings.Join(buffer, "")

	file.WriteString(outputStr)

}
