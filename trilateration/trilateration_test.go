package trilateration

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
)

var a = Coordinate{
	x: -3,
	y: -2,
	d: 7.21,
}
var b = Coordinate{
	x: 2,
	y: -2,
	d: 6.08,
}
var c = Coordinate{
	x: 3,
	y: 3,
	d: 2.23,
}

func TestTrilateration(t *testing.T) {
	fmt.Println("TestTrilateration")
	x, y := a.Trilaterate(b, c)
	fmt.Println("Coordinates of intersections center using formula x:", x, "y:", y)
	x1, y1 := b.Trilaterate(c, a)
	fmt.Println("b->(c,a) x:", x1, "y:", y1)
	x2, y2 := c.Trilaterate(b, a)
	fmt.Println("c->(b,a) x:", x2, "y:", y2)
}

func TestGetLocation(t *testing.T) {
	fmt.Println("TestGetLocation")
	x, y := GetLocation(7.21, 6.08, 2.23)
	//fmt.Println(GetLocation(7.21, 6.08, 2.23))
	fmt.Printf("x: %.1f y: %.1f \n", x, y)
}

func TestGetLocation2(t *testing.T) {
	fmt.Println("TestGetLocation2")
	//ejemplo graficamente real para punto D -100,400
	x, y := GetLocation2(721.11, 538.51, 670.82)
	fmt.Printf("x: %.1f y: %.1f \n", x, y)
	if auxX, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", x), 32); auxX != -100 {
		t.Errorf("TestGetLocation FAILED: Expected x = -100 && y = 400 and got x=%.2f", auxX)
	} else {
		t.Log("TestGetLocation X validation PASSED")
	}
	if auxY, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", y), 32); auxY != 400 {
		t.Errorf("TestGetLocation FAILED: Expected y = 400 and got y=%.2f", auxY)
	} else {
		t.Log("TestGetLocation Y validation PASSED")
	}

}

/*******challa********/

func TestType(t *testing.T) {
	fmt.Println(reflect.TypeOf("tst"))
	if math.IsNaN(math.Log(1.0)) {
		fmt.Println("nan")
	} else {
		fmt.Println("not nan")
	}
}

func TestReadCoorValues(t *testing.T) {
	x := GetAvgCoorValue(1.0017699999999998, 3.9989959999999996, math.NaN())
	fmt.Println(fmt.Sprintf("%.2f", x))

}

func TestHello(t *testing.T) {
	fmt.Println(Hello("Bastian"))
}
