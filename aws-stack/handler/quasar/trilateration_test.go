package main

import (
	"fmt"
	"math"
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
		t.Errorf("TestGetLocation FAILED: Expected x = -100 and got x=%.2f", auxX)
	}
	if auxY, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", y), 32); auxY != 400 {
		t.Errorf("TestGetLocation FAILED: Expected y = 400 and got y=%.2f", auxY)
	}
}

// TestAvgCoorValue percentage test of 3 values where one is nan
func TestAvgCoorValue(t *testing.T) {
	if x := GetAvgCoorValue(1.0017699999999998, 3.9989959999999996, math.NaN()); x != float32(2.500383) {
		t.Errorf("TestAvgCoorValue FAILED: Expected x = 2.5 and got x=%f", x)
	}
}
