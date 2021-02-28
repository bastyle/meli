package trilateration

import (
	"fmt"
	"math"
)

type Coordinate struct {
	x    float64
	y    float64
	d    float64
	uuid string
}

func (c1 Coordinate) Trilaterate(c2, c3 Coordinate) (x, y float64) {
	return Trilaterate(c1, c2, c3)
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Trilaterate(c1, c2, c3 Coordinate) (x, y float64) {
	d1, d2, d3, i1, i2, i3, j1, j2, j3 := &c1.d, &c2.d, &c3.d, &c1.x, &c2.x, &c3.x, &c1.y, &c2.y, &c3.y
	x = ((((math.Pow(*d1, 2)-math.Pow(*d2, 2))+(math.Pow(*i2, 2)-math.Pow(*i1, 2))+(math.Pow(*j2, 2)-math.Pow(*j1, 2)))*(2**j3-2**j2) - ((math.Pow(*d2, 2)-math.Pow(*d3, 2))+(math.Pow(*i3, 2)-math.Pow(*i2, 2))+(math.Pow(*j3, 2)-math.Pow(*j2, 2)))*(2**j2-2**j1)) / ((2**i2-2**i3)*(2**j2-2**j1) - (2**i1-2**i2)*(2**j3-2**j2)))
	y = ((math.Pow(*d1, 2) - math.Pow(*d2, 2)) + (math.Pow(*i2, 2) - math.Pow(*i1, 2)) + (math.Pow(*j2, 2) - math.Pow(*j1, 2)) + x*(2**i1-2**i2)) / (2**j2 - 2**j1)
	return x, y
}

/////////versión float 32
/*type Coordinate32 struct {
	x    float32
	y    float32
	d    float32
	uuid string
}

func (c1 Coordinate32) Trilaterate32(c2, c3 Coordinate32) (x, y float32) {
	return Trilaterate(c1, c2, c3)
}*/

func GetLocation(distances ...float32) (x, y float32) {
	fmt.Println("GetLocation", distances)
	var kenobi = Coordinate{
		x: -3,
		y: -2,
		//d: 7.21,
		d: float64(distances[0]),
	}
	var skywalker = Coordinate{
		x: 2,
		y: -2,
		//d: 6.08,
		d: float64(distances[1]),
	}
	var sato = Coordinate{
		x: 3,
		y: 3,
		//d: 2.23,
		d: float64(distances[2]),
	}
	imperioX, imperioY := kenobi.Trilaterate(skywalker, sato)
	return float32(imperioX), float32(imperioY)
	//return 1.0017699999999998, 3.9989959999999996
}
