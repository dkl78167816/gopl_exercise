// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			fill := getColorByHeight(i, j)
			r, g, b, _ := fill.RGBA()
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: #%2.2X%2.2X%2.2X'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, uint8(r), uint8(g), uint8(b))
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)

	if !ok {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, true

}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	res := math.Sin(r) / r
	if math.IsNaN(res) {
		return 0, false
	}
	return res, true
}

// ref: https://github.com/ivanbeldad/the-go-programming-language/
// treats height as 3 types, and generates 3 types of colors with bit operation
func getColorByHeight(i, j int) color.Color {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, _ := f(x, y)
	// -0.5 < sin(r)/r < 1
	// HH HH HH -> 00000000 00000000 00000000
	// -1 -> 0 -> 1
	// -1 = 00000000 00000000 11111111
	// 0 = 00000000 11111111 00000000
	// 1 = 11111111 00000000 00000000
	// total shift 16
	shift := uint8(math.Round((z + 1) * 8))
	var customColor uint32 = 0xFF
	customColor = customColor << shift
	b := uint8(customColor)
	customColor = customColor >> 8
	g := uint8(customColor)
	customColor = customColor >> 8
	r := uint8(customColor)
	// fmt.Printf("Color: %b\tR: %b\tG: %b\tB: %b\n", originalColor, r, g, b)
	// fmt.Printf("Shift: %d\n", shift)
	return color.RGBA{r, g, b, 0xFF}
}
