package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetFolderNames(t *testing.T) {
	names := []string{"pes", "fifa", "gta", "pes(2019)"}
	expected := []string{"pes", "fifa", "gta", "pes(2019)"}
	actual := GetFolderNames(names)
	assert.DeepEqual(t, expected, actual)
}

func TestGetFolderNames2(t *testing.T) {
	names := []string{"gta", "gta(1)", "gta", "avalon"}
	expected := []string{"gta", "gta(1)", "gta(2)", "avalon"}
	actual := GetFolderNames(names)
	assert.DeepEqual(t, expected, actual)
}

func TestGetFolderNames3(t *testing.T) {
	names := []string{"wano", "wano", "wano", "wano"}
	expected := []string{"wano", "wano(1)", "wano(2)", "wano(3)"}
	actual := GetFolderNames(names)
	assert.DeepEqual(t, expected, actual)
}

func TestGetFolderNames4(t *testing.T) {
	names := []string{"kaido", "kaido(1)", "kaido", "kaido(1)"}
	expected := []string{"kaido", "kaido(1)", "kaido(2)", "kaido(1)(1)"}
	actual := GetFolderNames(names)
	assert.DeepEqual(t, expected, actual)
}
