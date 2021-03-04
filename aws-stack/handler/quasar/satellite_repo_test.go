package main

import "testing"

func TestUpdateSatellite(t *testing.T) {
	t.Log("TestUpdateSatellite:::: ")
	if err := UpdateSatellite(11.5, "kenobi"); err != nil {
		t.Errorf("Error UpdateSatellite = %v.\n", err)
	} else {
		t.Log("la hiciste:::: ")
	}
}
