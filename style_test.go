package tint_test

import (
	"fmt"

	. "github.com/jwaldrip/tint"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Style", func() {

	var style *Style

	BeforeEach(func() {
		style = &Style{}
	})

	AfterEach(func() {
		fmt.Print(style.Render("#"))
	})

	Describe("Render", func() {
		It("should render unformated", func() {
			output := style.Render("*****")
			Expect(output).To(Equal("*****"))
		})

		Context("colors", func() {
			It("should render Black", func() {
				style.Color = Black
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("30m"))
			})

			It("should render Red", func() {
				style.Color = Red
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("31m"))
			})

			It("should render Green", func() {
				style.Color = Green
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("32"))
			})

			It("should render Yellow", func() {
				style.Color = Yellow
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("33m"))
			})

			It("should render Blue", func() {
				style.Color = Blue
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("34m"))
			})

			It("should render Magenta", func() {
				style.Color = Magenta
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("35m"))
			})

			It("should render Cyan", func() {
				style.Color = Cyan
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("36m"))
			})

			It("should render LightGrey", func() {
				style.Color = LightGrey
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("37m"))
			})

			It("should render LightBlack", func() {
				style.Color = LightBlack
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("90m"))
			})

			It("should render LightRed", func() {
				style.Color = LightRed
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("91m"))
			})

			It("should render LightGreen", func() {
				style.Color = LightGreen
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("92m"))
			})

			It("should render LightYellow", func() {
				style.Color = LightYellow
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("93m"))
			})

			It("should render LightBlue", func() {
				style.Color = LightBlue
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("94m"))
			})

			It("should render LightMagenta", func() {
				style.Color = LightMagenta
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("95m"))
			})

			It("should render LightCyan", func() {
				style.Color = LightCyan
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("96m"))
			})

			It("should render White", func() {
				style.Color = White
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("97m"))
			})

		})

		Context("backgrounds", func() {
			It("should render Black", func() {
				style.Background = Black
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("40m"))
			})

			It("should render Red", func() {
				style.Background = Red
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("41m"))
			})

			It("should render Green", func() {
				style.Background = Green
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("42m"))
			})

			It("should render Yellow", func() {
				style.Background = Yellow
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("43m"))
			})

			It("should render Blue", func() {
				style.Background = Blue
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("44m"))
			})

			It("should render Magenta", func() {
				style.Background = Magenta
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("45m"))
			})

			It("should render Cyan", func() {
				style.Background = Cyan
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("46m"))
			})

			It("should render LightGrey", func() {
				style.Background = LightGrey
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("47m"))
			})

			It("should render LightBlack", func() {
				style.Background = LightBlack
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("100m"))
			})

			It("should render LightRed", func() {
				style.Background = LightRed
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("101m"))
			})

			It("should render LightGreen", func() {
				style.Background = LightGreen
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("102m"))
			})

			It("should render LightYellow", func() {
				style.Background = LightYellow
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("103m"))
			})

			It("should render LightBlue", func() {
				style.Background = LightBlue
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("104m"))
			})

			It("should render LightMagenta", func() {
				style.Background = LightMagenta
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("105m"))
			})

			It("should render LightCyan", func() {
				style.Background = LightCyan
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("106m"))
			})

			It("should render White", func() {
				style.Background = White
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("107m"))
			})

		})

		Context("bold", func() {
			It("should render boldly", func() {
				style.Bold = true
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("1m"))
			})
		})

		Context("dim", func() {
			It("should render dimly", func() {
				style.Dim = true
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("2m"))
			})
		})

		Context("underline", func() {
			It("should render underline", func() {
				style.Underline = true
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("4m"))
			})
		})

		Context("blink", func() {
			It("should render underline", func() {
				style.Blink = true
				output := style.Render("*****")
				Expect(output).To(ContainSubstring("5m"))
			})
		})

	})

})
