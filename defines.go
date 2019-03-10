package main

import "encoding/xml"

type btsPos struct {
	x float64
	y float64
}

type Graphs struct {
	XMLName xml.Name `xml:"graphs"`
	Graphs  []Graph  `xml:"graph"`
}

type Graph struct {
	XMLName xml.Name `xml:"graph"`
	Name    string   `xml:"name"`
	Nodes   []Node   `xml:"node"`
	Edges   []Edge   `xml:"edge"`
}

type Node struct {
	XMLName xml.Name `xml:"node"`
	ID      string   `xml:"id"`
	Name    string   `xml:"name"`
}

type Edge struct {
	XMLName xml.Name `xml:"edge"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name"`
	From    string   `xml:"from,attr"`
	To      string   `xml:"to,attr"`
}
