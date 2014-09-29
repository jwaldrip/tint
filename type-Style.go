package tint

import "fmt"
import "strings"
import "strconv"

// Style is the struct that holds information on how to style a string
type Style struct {
	Color      int
	Background int
	Bold       bool
	Blink      bool
	Underline  bool
	Inverted   bool
	Use256     bool
}

// Render renders a style on a string
func (s *Style) Render(str string) string {
	return strings.Join([]string{s.escapeSequence(), str, s.resetSequence()}, "")
}

func (s *Style) getColor() string {
	clr := s.Color
	if clr == 0 && !s.Use256 {
		clr = DefaultColor
	}
	return strconv.Itoa(clr)
}

func (s *Style) getBackground() string {
	bg := s.Background
	if bg == 0 && !s.Use256 {
		bg = DefaultColor
	}
	return strconv.Itoa(bg + 10)
}

func (s *Style) resetSequence() string {
	return fmt.Sprintf("%s[%sm", escape, strconv.Itoa(reset))
}

func (s *Style) escapeSequence() string {
	var strs []string
	if s.Bold {
		strs = append(strs, strconv.Itoa(bold))
	}
	if s.Blink {
		strs = append(strs, strconv.Itoa(blink))
	}
	if s.Underline {
		strs = append(strs, strconv.Itoa(underline))
	}
	if s.Inverted {
		strs = append(strs, strconv.Itoa(invert))
	}
	if s.Use256 {
		strs = append(strs, "38;5")
	}
	strs = append(strs, s.getColor())
	if s.Use256 {
		strs = append(strs, "48;5")
	}
	strs = append(strs, s.getBackground())
	return fmt.Sprintf("%s[%sm", escape, strings.Join(strs, ";"))
}
