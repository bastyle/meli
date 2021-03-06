package main

import (
	"testing"
)

func TestIsThereNecessaryInfo(t *testing.T) {
	if items, err := GetAllDataSatell(); err != nil {
		t.Errorf("Error GetAllDataSatell = %v.\n", err)
	} else {
		if isOk, requestBody := IsThereNecessaryInfo(items); !isOk {
			t.Errorf("Error IsThereNecessaryInfo = %v.\n", isOk)
		} else {
			t.Log("requestBody: ", requestBody)
		}
	}
}
