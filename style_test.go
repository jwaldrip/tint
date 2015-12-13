package tint_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jwaldrip/tint"
)

var (
	input string = "*****"
)

func TestRenderColors(t *testing.T) {
	var tests = []struct {
		containsColor string
		color         int
		colorStr      string
	}{
		{"*****", -1, "unformatted"},
		{"30m", tint.Black, "black"},
		{"31m", tint.Red, "red"},
		{"32m", tint.Green, "green"},
		{"33m", tint.Yellow, "yellow"},
		{"34m", tint.Blue, "blue"},
		{"35m", tint.Magenta, "magenta"},
		{"36m", tint.Cyan, "cyan"},
		{"37m", tint.LightGrey, "light grey"},
		{"90m", tint.LightBlack, "light black"},
		{"91m", tint.LightRed, "light red"},
		{"92m", tint.LightGreen, "light green"},
		{"93m", tint.LightYellow, "light yellow"},
		{"94m", tint.LightBlue, "light blue"},
		{"95m", tint.LightMagenta, "light magenta"},
		{"96m", tint.LightCyan, "light cyan"},
		{"97m", tint.White, "white"},
	}
	for _, test := range tests {
		style := &tint.Style{}
		if test.color != -1 {
			style.Color = test.color
		}
		output := style.Render(input)
		if !strings.Contains(output, test.containsColor) {
			fail := fmt.Sprintf("it should render the color %s", test.colorStr)
			t.Error(fail)
		}
		if !strings.Contains(output, input) {
			t.Error("it should contain the orignal string")
		}
	}

}

func TestStyleBackgrounds(t *testing.T) {
	var tests = []struct {
		containsColor string
		color         int
		colorStr      string
	}{
		{"40m", tint.Black, "black"},
		{"41m", tint.Red, "red"},
		{"42m", tint.Green, "green"},
		{"43m", tint.Yellow, "yellow"},
		{"44m", tint.Blue, "blue"},
		{"45m", tint.Magenta, "magenta"},
		{"46m", tint.Cyan, "cyan"},
		{"47m", tint.LightGrey, "light grey"},
		{"100m", tint.LightBlack, "light black"},
		{"101m", tint.LightRed, "light red"},
		{"102m", tint.LightGreen, "light green"},
		{"103m", tint.LightYellow, "light yellow"},
		{"104m", tint.LightBlue, "light blue"},
		{"105m", tint.LightMagenta, "light magenta"},
		{"106m", tint.LightCyan, "light cyan"},
		{"107m", tint.White, "white"},
	}
	for _, test := range tests {
		style := &tint.Style{}
		style.Background = test.color
		output := style.Render(input)
		if !strings.Contains(output, test.containsColor) {
			fail := fmt.Sprintf("it should render the background color %s", test.colorStr)
			t.Error(fail)
		}
		if !strings.Contains(output, input) {
			t.Error("it should contain the orignal string")
		}
	}
}

func TestStyleBold(t *testing.T) {
	style := &tint.Style{
		Bold: true,
	}
	output := style.Render(input)
	if !strings.Contains(output, input) {
		t.Error("it should contain the orignal string")
	}
	if !strings.Contains(output, "1m") {
		t.Error("it should have rendered the text as bold")
	}
}

func TestStyleDim(t *testing.T) {
	style := &tint.Style{
		Dim: true,
	}
	output := style.Render(input)
	if !strings.Contains(output, input) {
		t.Error("it should contain the orignal string")
	}
	if !strings.Contains(output, "2m") {
		t.Error("it should have rendered the text as dim")
	}
}

func TestStyleUnderline(t *testing.T) {
	style := &tint.Style{
		Underline: true,
	}
	output := style.Render(input)
	if !strings.Contains(output, input) {
		t.Error("it should contain the orignal string")
	}
	if !strings.Contains(output, "4m") {
		t.Error("it should have rendered the text with an underline")
	}
}

func TestStyleBlink(t *testing.T) {
	style := &tint.Style{
		Blink: true,
	}
	output := style.Render(input)
	if !strings.Contains(output, input) {
		t.Error("it should contain the orignal string")
	}
	if !strings.Contains(output, "5m") {
		t.Error("it should have rendered the text blinking")
	}
}
