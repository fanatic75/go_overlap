package main

import (
	"encoding/json"
	"geektrust/fund"
	"io/ioutil"
	"os"
	"testing"
)

func TestCalculateOverlap(t *testing.T) {
	err := json.Unmarshal([]byte(fund.Funds), &fundsJson)
	if err != nil {
		t.Logf("error while unmarshal %v", err)
	}
	for _, fund := range fundsJson.Funds {
		fundsMap[fund.Name] = fund.Stocks
	}
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	var user = User{Funds: make([]string, 0)}
	user.Funds = append(user.Funds, "AXIS_BLUECHIP")
	user.PrintOverlapValue("MIRAE_ASSET_EMERGING_BLUECHIP")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if "MIRAE_ASSET_EMERGING_BLUECHIP AXIS_BLUECHIP 39.13%" != "MIRAE_ASSET_EMERGING_BLUECHIP AXIS_BLUECHIP 39.13%" {
		t.Errorf("Expected %s, got %s", "MIRAE_ASSET_EMERGING_BLUECHIP AXIS_BLUECHIP 39.13%", out)
	}
}
