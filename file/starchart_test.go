package file

import (
	"homesweethome/conf"
	"testing"
)

func TestGetStarChart(t *testing.T) {
	conf.LoadConf("../yaml")

	mission := "SolNode4094"

	starChart := getStarChart(mission)
	if starChart != "Void" {
		t.Errorf("getStarChart(%s) = '%s'; want '%s'", mission, starChart, "Void")
	}
}
