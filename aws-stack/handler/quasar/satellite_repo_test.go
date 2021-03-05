package main

import (
	"fmt"
	"testing"
)

func TestUpdateDistanceSatellite(t *testing.T) {
	if err := UpdateDistanceSatellite(11.5, "kenobi"); err != nil {
		t.Errorf("Error UpdateSatellite = %v.\n", err)
	}
}

/*func TestGetDistance(t *testing.T) {
	if err := GetDistance("kenobi"); err != nil {
		t.Errorf("Error GetDistance = %v.\n", err)
	}
}*/

/*func TestGetList(t *testing.T) {
	var exampleMsg = [5]string{"esto", "", "un", "mensaje", ""}
	if msg, err := GetList(exampleMsg[:]); err != nil {
		t.Errorf("Error GetList = %v.\n", err)
	} else {
		t.Log("msg: ", msg)
	}
}*/

func TestGetDataSatell(t *testing.T) {
	t.Log("TestSatellites:")
	//item := DataSat{}
	if item, err := GetDataSatell("sato"); err != nil {
		t.Errorf("Error get Data = %v.\n", err)
	} else {
		t.Log("item:", item)
		fmt.Printf("len msg: %v.\n", len(item.Message))
	}
}

func TestUpdateSatellite(t *testing.T) {
	msg := [5]string{"kenobi", "", "un", "mensaje", ""}
	//t.Log("arreglo?????:", msg[:])
	if err := UpdateSatellite("kenobi", 314.6, msg[:]); err != nil {
		t.Errorf("Error UpdateSatellite = %v.\n", err)
	} else {
		t.Log("satellite updated.")
	}
}
