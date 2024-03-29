package main

import (
	"fmt"
	"math"
)

// Satellite struct that represent a fisical satellite.
type Satellite struct {
	x    float64
	y    float64
	d    float64
	name string
}

var kenobiSat = Satellite{
	x: -500,
	y: -200,
	//d:    7.21,
	name: "kenobi",
}
var skywalkerSat = Satellite{
	x: 100,
	y: -100,
	//d:    6.08,
	name: "skywalker",
}
var satoSat = Satellite{
	x: 500,
	y: 100,
	//d:    2.23,
	name: "sato",
}

// GetAvgCoorValue get avg from a series on values delivered.
func GetAvgCoorValue(coordinateValues ...float64) float32 {
	var finalCoor float32
	var coorSum float64
	var coorCount float64
	for _, v := range coordinateValues {
		if !math.IsNaN(v) && !math.IsInf(v, 0) {
			coorCount += 1
			coorSum += v
		}
	}
	finalCoor = float32(coorSum / coorCount)
	return finalCoor
}

// Trilaterate obtains the trilation (location x, y) given 3 known points and their distances to a fourth unknown point.
func (c1 Satellite) Trilaterate(c2, c3 Satellite) (x, y float64) {
	return Trilaterate(c1, c2, c3)
}

// Trilaterate obtains the trilation (location x, y) given 3 known points and their distances to a fourth unknown point.
func Trilaterate(c1, c2, c3 Satellite) (x, y float64) {
	d1, d2, d3, i1, i2, i3, j1, j2, j3 := &c1.d, &c2.d, &c3.d, &c1.x, &c2.x, &c3.x, &c1.y, &c2.y, &c3.y
	x = ((((math.Pow(*d1, 2)-math.Pow(*d2, 2))+(math.Pow(*i2, 2)-math.Pow(*i1, 2))+(math.Pow(*j2, 2)-math.Pow(*j1, 2)))*(2**j3-2**j2) - ((math.Pow(*d2, 2)-math.Pow(*d3, 2))+(math.Pow(*i3, 2)-math.Pow(*i2, 2))+(math.Pow(*j3, 2)-math.Pow(*j2, 2)))*(2**j2-2**j1)) / ((2**i2-2**i3)*(2**j2-2**j1) - (2**i1-2**i2)*(2**j3-2**j2)))
	y = ((math.Pow(*d1, 2) - math.Pow(*d2, 2)) + (math.Pow(*i2, 2) - math.Pow(*i1, 2)) + (math.Pow(*j2, 2) - math.Pow(*j1, 2)) + x*(2**i1-2**i2)) / (2**j2 - 2**j1)
	return x, y
}

// GetLocation obtains the location of unknown point (x, y) given 3 known distances to it.
func GetLocation(distances ...float32) (x, y float32) {
	fmt.Println("GetLocation distances: ", distances)

	kenobiSat.d = float64(distances[0])
	skywalkerSat.d = float64(distances[1])
	satoSat.d = float64(distances[2])

	xFromKenobi, yFromKenobi := kenobiSat.Trilaterate(skywalkerSat, satoSat)
	xFromSkywalker, yFromSkywalker := skywalkerSat.Trilaterate(satoSat, kenobiSat)
	xFromSato, yFromSato := satoSat.Trilaterate(skywalkerSat, kenobiSat)

	return GetAvgCoorValue(xFromKenobi, xFromSato, xFromSkywalker), GetAvgCoorValue(yFromKenobi, yFromSato, yFromSkywalker)
}
