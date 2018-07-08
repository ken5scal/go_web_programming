package main

import (
	"image"
	"image/color"
)

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

func resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	ratio := bounds.Dx() / newWidth
	out := image.NewNRGBA(image.Rect(
		bounds.Min.X/ratio,
		bounds.Min.Y/ratio,
		bounds.Max.X/ratio,
		bounds.Max.Y/ratio))
	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1{
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, i+1{
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)})
		}
	}
	return *out
}