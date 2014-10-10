package tint_test

import (
	"fmt"

	. "github.com/jwaldrip/tint"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var output string

var _ = Describe("functions", func() {

	BeforeEach(func() {
		output = ""
	})

	AfterEach(func() {
		fmt.Print(output)
	})

	Describe("Stylize", func() {
		It("should stylize a string", func() {
			output = Stylize("#", Style{Color: Blue, Underline: true})
			Expect(output).To(ContainSubstring("4;34m"))
		})
	})

	Describe("Colorize", func() {
		It("should colorize a string", func() {
			output = Colorize("#", Blue)
			Expect(output).To(ContainSubstring("34m"))
		})
	})

	Describe("ColorizeBackground", func() {
		It("should colorize a background of a string", func() {
			output = ColorizeBackground("#", Blue)
			Expect(output).To(ContainSubstring("44m"))
		})
	})
})
