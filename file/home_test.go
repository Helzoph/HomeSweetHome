package file

import (
	"homesweethome/conf"
	"testing"
)

func TestGetHome(t *testing.T) {
	conf.LoadConf("../yaml")

	str := "6581.074 Sys [Info]: I: /Lotus/Levels/Orokin/LargeTieredIntermediate.level"
	starChart := "Void"
	home := make([]string, 0)

	getHome(str, starChart, &home)
	if len(home) <= 0 {
		t.Errorf("getHome('%s', '%s', &home) = '%v'; want '%v'", str, starChart, home, "CorpusGasCityRemaster")
	}
}
