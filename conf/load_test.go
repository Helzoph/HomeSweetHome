package conf

import (
	"testing"
)

func TestLoadConf(t *testing.T) {
	LoadConf("../yaml")

	if len(StarChartList) == 0 {
		t.Error("StarChartList is empty")
	}

	if len(Missions) == 0 {
		t.Error("Missions is empty")
	}

	if len(StarChartMap) == 0 {
		t.Error("StarChartMap is empty")
	}
}
