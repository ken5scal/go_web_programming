package main

import (
	"image"
	"fmt"
	"os"
	"net/http"
	"html/template"
	"time"
	"strconv"
	"image/draw"
	"bytes"
	"image/jpeg"
	"encoding/base64"
	"sync"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	// building up the source tile database
	TILESDB = tilesDB()
	fmt.Println("Mosaic server started.")
	server.ListenAndServe()
}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	// get the content from the POSTed form
	r.ParseMultipartForm(10485760) // max body in memory is 10MB
	file, _, _ := r.FormFile("image")
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))
	// decode and get original image
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	// build up the tiles database
	db := cloneTilesDB()

	// Fanning out, cutting up image for independent processing
	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
	c2 := cut(original, &db, tileSize, bounds.Min.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := cut(original, &db, tileSize, bounds.Min.X/2, bounds.Min.Y/2, bounds.Max.X, bounds.Max.Y)
	c := combine(bounds, c1, c2, c3, c4)

	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	t1 := time.Now()
	images := map[string]string {
		"original": originalStr,
		"mosaic": <- c,
		"duration": fmt.Sprint("%v ", t1.Sub(t0)),
	}

	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)
}

func cut(original image.Image, db *DB, tileSize, x1, y1, x2, y2 int) <-chan image.Image {
	c := make(chan image.Image)
	sp := image.Point{0, 0}
	go func() {
		// create a new image for the mosaic (IN CASE OF Processing entire image at once)
		newimage := image.NewNRGBA(image.Rect(x1, y1, x2, y2))
		for y := y1; y < y2; y = y + tileSize {
			for x := x1; x < x2; x = x + tileSize {
				// use the top left most pixel as the average color
				r, g, b, _ := original.At(x, y).RGBA()
				color := [3]float64{float64(r), float64(g), float64(b)}
				// get the closest tile from the tiles DB
				nearest := db.nearest(color)
				file, err := os.Open(nearest)
				if err == nil {
					img, _, err := image.Decode(file)
					if err == nil {
						// resize the tile to the correct size
						t := resize(img, tileSize)
						tile := t.SubImage(t.Bounds())
						tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
						// draw the tile into the mosaic
						draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
					} else {
						fmt.Println("error:", err, nearest)
					}
				} else {
					fmt.Println("error:", nearest)
				}
				file.Close()
			}
		}
		c <- newimage.SubImage(newimage.Rect)
	}()

	return c
}

// returning a receive-only channel
func combine(r image.Rectangle, c1,c2,c3,c4 <- chan image.Image) <- chan string {
	c := make(chan string)

	go func() {
		// wait until all subimages copied to final image
		var wg sync.WaitGroup
		img := image.NewNRGBA(r)

		copy := func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
			draw.Draw(dst, r, src, sp, draw.Src)
			// decrements counter as subimages copied over
			wg.Done()
		}

		wg.Add(4)
		var s1, s2, s3, s4 image.Image
		var ok1, ok2, ok3, ok4 bool

		for {
			// start processing whichever segment that comes first
			select {
			case s1, ok1 = <- c1:
				go copy(img, s1.Bounds(), s1, image.Point{r.Min.X, r.Min.Y})
			case s2, ok2 = <- c2:
				go copy(img, s2.Bounds(), s2, image.Point{r.Max.X/2, r.Min.Y})
			case s2, ok2 = <- c3:
				go copy(img, s2.Bounds(), s3, image.Point{r.Min.X, r.Max.Y/2})
			case s2, ok2 = <- c4:
				go copy(img, s4.Bounds(), s4, image.Point{r.Max.X/2, r.Max.Y/2})
			}
			if (ok1 && ok2 && ok3 && ok4) {
				break
			}
		}

		wg.Wait()
		buf2 := new(bytes.Buffer)
		jpeg.Encode(buf2, img, nil)
		c <- base64.StdEncoding.EncodeToString(buf2.Bytes())
	}()

	return c
}