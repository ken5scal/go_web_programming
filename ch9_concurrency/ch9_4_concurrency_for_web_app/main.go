package main

import (
	"image"
	"image/color"
	"fmt"
	"io/ioutil"
	"os"
	"math"
	"net/http"
	"html/template"
	"time"
	"strconv"
	"image/draw"
	"bytes"
	"image/jpeg"
	"encoding/base64"
)

var TILESDB map[string][3]float64

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr : "127.0.0.1:8080",
		Handler: mux,
	}

	TILESDB = titleDB()
	server.ListenAndServe()
}
func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	r.ParseMultipartForm(10485760)
	file, _, _ := r.FormFile("image")
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))

	original, _, _ := image.Decode(file)
	bounds := original.Bounds()

	newImage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))

	// every time image uploaded, clone the db
	db := cloneTilesDB()

	// we take the top-left pixel and assume that's the average color for each tile
	sp := image.Point{0, 0}

	// Iterates through target image
	for y:= bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {
			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{float64(r), float64(g), float64(b)}

			nearest := nearest(color, &db)
			file, err := os.Open(nearest)
			if err == nil {
				img, _, err := image.Decode(file)
				if err == nil {
					t := resize(img, tileSize)
					tile := t.SubImage(t.Bounds())
					tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
					draw.Draw(newImage, tileBounds, tile, sp, draw.Src)
				} else {
					fmt.Println("error: ", err, nearest)
				}
			} else {
				fmt.Println("error: ", nearest)
			}
			file.Close()
		}
	}

	// Encode in JPEG, deliver to browser in base64 string
	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, newImage, nil)
	mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())
	t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic": mosaic,
		"duration": fmt.Sprintf("%v ", t1.Sub(t0)),
	}
	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)
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
	for y, new_y := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, new_y = y+ratio, new_y+1{
		for x, new_x := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, new_x = x+ratio, new_x+1{
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(new_x, new_y, color.NRGBA{uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)})
		}
	}
	return *out
}

// creates a db of the tile picture by scanning the dir
// it will store average color.
func titleDB() map[string][3]float64 {
	fmt.Println("Start populating titles db ...")
	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("titles")
	for _, f := range files {
		name := "tiles/" + f.Name()
		file, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating TILEDB:", err, name)
			}
		} else {
			fmt.Println("cannot open file", name ,err)
		}
		file.Close()
	}
	fmt.Println("Finished populating tiles db.")
	return db
}

// finding the nearest matching image
func nearest(target [3]float64, db *map[string][3]float64) string {
	var filename string
	smallest := 1000000.0
	for k, v := range *db {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(*db, filename)
	return filename
}

func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1])+sq(p2[2]-p1[2]))
}

func sq(n float64) float64 {
	return n * n
}

// clone the tile db each tiem the photo mosaic is generated
func cloneTilesDB() map[string][3]float64 {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] =v
	}
	return db
}
