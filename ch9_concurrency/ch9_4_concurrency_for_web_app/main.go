package main

import "image"

func main() {
}

func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r_tot, g_tot, b_tot := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r_tot, g_tot, b_tot = r_tot+float64(r1), g_tot+float64(g1), b_tot+float64(b1)
		}
	}
	totalNumberOfPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r_tot / totalNumberOfPixels, g_tot / totalNumberOfPixels, b_tot / totalNumberOfPixels}
}