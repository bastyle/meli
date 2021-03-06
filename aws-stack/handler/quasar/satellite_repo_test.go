package main

import (
	"fmt"
	"testing"
)

func TestGetDataSatell(t *testing.T) {
	t.Log("TestSatellites:")
	exampleName := "sato"
	if item, err := GetDataSatell(exampleName); err != nil {
		t.Errorf("Error get Data = %v.\n", err)
	} else if item.Name == exampleName {
		t.Log("item:", item)
		fmt.Printf("len msg: %v.\n", len(item.Message))
	} else {
		t.Log("no existe satelite buscado.")
	}
}

func TestGetAllDataSatell(t *testing.T) {
	if items, err := GetAllDataSatell(); err != nil {
		t.Errorf("Error GetAllDataSatell = %v.\n", err)
	} else {
		t.Log("item.X:: ", items[0].X)
		t.Log("item.X:: ", items[1].X)
		t.Log("item.X:: ", items[2].X)
	}
}

/*func TestUpdateSatellite(t *testing.T) {
	msg := [1]string{""}
	if err := UpdateSatellite("kenobi", 100, msg[:]); err != nil {
		t.Errorf("Error UpdateSatellite = %v.\n", err)
	} else {
		t.Log("satellite updated.")
	}
}*/

func TestUpdateFakeSatellite(t *testing.T) {
	msg := [1]string{""}
	if err := UpdateSatellite("fakeSatellite", 100, msg[:]); err != nil {
		t.Log("satellite dosnt exists.")
	} else {
		t.Errorf("Error UpdateSatellite = %v.\n", err)
	}
}

/*func TestResetSatellDynamicData(t *testing.T) {
	if err := ResetSatellDynamicData(); err != nil {
		t.Errorf("Error ResetSatellDynamicData = %v.\n", err)
	} else {
		t.Log("satellites have been updated.")
	}
}*/
