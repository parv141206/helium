// Hello there, parv here!
// This is a simple project to make ASCII art from JPG images.
// Currently it only works with JPGs :(
// But well
// I have documented this as much as i could, its simple, still
// Enjoy :)

package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"math"
	"os"
)

func main() {
	path := ""
	println("Please enter the path to image: ")
	fmt.Scan(&path)

	// Usual file handling,

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// This gradient defines the characters used to make the image. Currently i have kept these many:
	gradient := "@%#*+=-:. "

	bounds := img.Bounds()

	// NOTE: The scale here shows how many pixels to skip. A better algorithm would be to find average i guess but currently i am using this "skipping" logic.
	// Basically it makes sure that the result is not overflowing and is of proper scale.
	scale := 30
	if bounds.Max.X > 3000 {
		scale = 30
	} else if bounds.Max.X > 2000 {
		scale = 15
	} else {
		scale = 5
	}

	// Following is the main loops for processing each pixel. The outer for Y direction and inner for X.
	for y := bounds.Min.Y; y < bounds.Max.Y; y += scale {
		for x := bounds.Min.X; x < bounds.Max.X; x += scale {
			pixelColor := img.At(x, y)
			r, g, b, _ := pixelColor.RGBA()

			red := float64(r / 257)
			green := float64(g / 257)
			blue := float64(b / 257)

			// The following luminance formula is somehow standard and since i dont have much better formula i am using this currently
			luminance := 0.2126*red + 0.7152*green + 0.0722*blue

			index := int(math.Round(luminance / 255 * float64(len(gradient)-1)))
			fmt.Print(string(gradient[index]), " ")
		}
		fmt.Println()
	}

}
