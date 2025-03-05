package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy) // Create a slice of slices
	
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx) // Create a slice for each row
		for x := 0; x < dx; x++ {
			row[x] = uint8(x * y)
		}
		image[y] = row // Assign the row to the image
	}
	
	return image
}

func main() {
	pic.Show(Pic)
}
