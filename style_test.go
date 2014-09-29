package tint_test

import (
	"fmt"
	"strconv"

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
		Context("colors", func() {
			It("should render Black", func() {
				style.Color = Black
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Black)))
			})

			It("should render Red", func() {
				style.Color = Red
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Red)))
			})

			It("should render Green", func() {
				style.Color = Green
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Green)))
			})

			It("should render Yellow", func() {
				style.Color = Yellow
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Yellow)))
			})

			It("should render Blue", func() {
				style.Color = Blue
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Blue)))
			})

			It("should render Magenta", func() {
				style.Color = Magenta
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Magenta)))
			})

			It("should render Cyan", func() {
				style.Color = Cyan
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Cyan)))
			})

			It("should render LightGrey", func() {
				style.Color = LightGrey
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightGrey)))
			})

			It("should render LightBlack", func() {
				style.Color = LightBlack
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightBlack)))
			})

			It("should render LightRed", func() {
				style.Color = LightRed
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightRed)))
			})

			It("should render LightGreen", func() {
				style.Color = LightGreen
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightGreen)))
			})

			It("should render LightYellow", func() {
				style.Color = LightYellow
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightYellow)))
			})

			It("should render LightBlue", func() {
				style.Color = LightBlue
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightBlue)))
			})

			It("should render LightMagenta", func() {
				style.Color = LightMagenta
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightMagenta)))
			})

			It("should render LightCyan", func() {
				style.Color = LightCyan
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightCyan)))
			})

			It("should render White", func() {
				style.Color = White
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(White)))
			})

		})

		Context("backgrounds", func() {
			It("should render Black", func() {
				style.Background = Black
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Black + 10)))
			})

			It("should render Red", func() {
				style.Background = Red
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Red + 10)))
			})

			It("should render Green", func() {
				style.Background = Green
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Green + 10)))
			})

			It("should render Yellow", func() {
				style.Background = Yellow
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Yellow + 10)))
			})

			It("should render Blue", func() {
				style.Background = Blue
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Blue + 10)))
			})

			It("should render Magenta", func() {
				style.Background = Magenta
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Magenta + 10)))
			})

			It("should render Cyan", func() {
				style.Background = Cyan
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(Cyan + 10)))
			})

			It("should render LightGrey", func() {
				style.Background = LightGrey
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightGrey + 10)))
			})

			It("should render LightBlack", func() {
				style.Background = LightBlack
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightBlack + 10)))
			})

			It("should render LightRed", func() {
				style.Background = LightRed
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightRed + 10)))
			})

			It("should render LightGreen", func() {
				style.Background = LightGreen
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightGreen + 10)))
			})

			It("should render LightYellow", func() {
				style.Background = LightYellow
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightYellow + 10)))
			})

			It("should render LightBlue", func() {
				style.Background = LightBlue
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightBlue + 10)))
			})

			It("should render LightMagenta", func() {
				style.Background = LightMagenta
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightMagenta + 10)))
			})

			It("should render LightCyan", func() {
				style.Background = LightCyan
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(LightCyan + 10)))
			})

			It("should render White", func() {
				style.Background = White
				output := style.Render("*****")
				Ω(output).To(ContainSubstring(strconv.Itoa(White + 10)))
			})

		})
	})

})
