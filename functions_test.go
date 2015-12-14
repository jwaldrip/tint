package tint_test

import (
	"strings"
	"testing"

	"github.com/jwaldrip/tint"
)

var output string

func TestStylize(t *testing.T) {
	output := tint.Stylize("#", tint.Style{Color: tint.Blue, Underline: true})
	if !strings.Contains(output, "4;34m") {
		t.Error("It should have colored the string blue with an underline.")
	}
}

func TestColorize(t *testing.T) {
	output := tint.Colorize("#", tint.Blue)
	if !strings.Contains(output, "34m") {
		t.Error("It should have colored the string blue.")
	}
}

func TestColorizeBackground(t *testing.T) {
	output := tint.ColorizeBackground("#", tint.Blue)
	if !strings.Contains(output, "44m") {
		t.Error("It should have given the string a blue background.")
	}
}
