package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func decode() {
	catFile, err := os.Open("./cat.png")
	if err != nil {
		log.Fatal(err)
	}
	defer catFile.Close()

	imData, imType, err := image.Decode(catFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(imData)
	fmt.Println(imType)

	cat, err := png.Decode(catFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cat)
}

func main() {
	catFile, err := os.Open("./cat.png")
	if err != nil {
		log.Fatal(err)
	}
	defer catFile.Close()

	img, err := png.Decode(catFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(img)

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {

		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}

		fmt.Print("\n")
	}
}
