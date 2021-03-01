package trilateration

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"
)

// TestGetLocation test of a graphically real example
func TestGetLocation(t *testing.T) {
	fmt.Println("TestGetLocation")
	//ejemplo graficamente real para punto D -100,400
	x, y := GetLocation(721.11, 538.51, 670.82)
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

// TestAvgCoorValue percentage test of 3 values where one is nan
func TestAvgCoorValue(t *testing.T) {
	if x := GetAvgCoorValue(1.0017699999999998, 3.9989959999999996, math.NaN()); x != float32(2.500383) {
		t.Errorf("TestAvgCoorValue FAILED: Expected x = 2.5 and got x=%f", x)
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

type Message struct {
	Name string
	Body string
	age  int
}

func TestHello(t *testing.T) {
	//fmt.Println(Hello("Bastian"))
	m := Message{"Alice", "Hello", 34}
	body, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("Bastian", body)
	}

}
