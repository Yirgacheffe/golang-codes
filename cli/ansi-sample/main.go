package main

import (
	"fmt"

	"github.com/go-playground/ansi/v3"
)

const blinkRed = ansi.Red + ansi.Blink + ansi.Underline
const textLine = "text line of the codes ..."

func main() {

	fmt.Printf("%s%s%s\n", ansi.Black+ansi.GrayBackground, textLine, ansi.Reset)

	// Font
	fmt.Printf("%s%s%s\n", ansi.Gray, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.DarkGray, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.Red, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.LightRed, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.Green, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.LightGreen, textLine, ansi.Reset)

	fmt.Printf("%s%s%s\n", ansi.Yellow, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.LightYellow, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.Blue, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.LightBlue, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.Magenta, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.LightMagenta, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.Cyan, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.LightCyan, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n\n", ansi.White, textLine, ansi.Reset)

	//
	fmt.Printf("%s%s%s\n", ansi.Gray+ansi.BlackBackground, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.Black+ansi.RedBackground, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.Black+ansi.GreenBackground, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.Black+ansi.YellowBackground, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.Black+ansi.BlueBackground, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.Black+ansi.MagentaBackground, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n", ansi.Black+ansi.CyanBackground, textLine, ansi.Reset)
	fmt.Printf("%s%s%s\n\n", ansi.Black+ansi.GrayBackground, textLine, ansi.Reset)

	fmt.Printf("%s%s%s\n\n", ansi.Inverse, textLine, ansi.InverseOff)

	fmt.Printf("%s%s%s\n\n", ansi.Italics, textLine, ansi.ItalicsOff)

	fmt.Printf("%s%s%s\n\n", ansi.Underline, "underline code ...", ansi.UnderlineOff)

	fmt.Printf("%s%s%s\n\n", ansi.Blink, "blink code line ...", ansi.BlinkOff)

	fmt.Printf("%s%s%s %s\n\n", ansi.Bold, "test bold", ansi.BoldOff, "and normal again!")

	fmt.Printf("%s%s%s\n", blinkRed, "blink red & underline", ansi.Reset)

}
