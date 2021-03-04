package main

import "testing"

func TestUpdateSatellite(t *testing.T) {
	t.Log("TestUpdateSatellite:::: ")
	if err := UpdateSatellite("kenobi", 115.5); err != nil {
		t.Errorf("Error UpdateSatellite = %v.\n", err)
	} else {
		t.Log("la hiciste:::: ")
	}
}
