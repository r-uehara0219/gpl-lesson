// Run with "web" command-line argument for web server.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		constants := setOptions(r.URL.Query())
		w.Header().Add("Content-Type", "image/svg+xml")
		surface(w, constants)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type constants struct {
	sin30   float64
	cos30   float64
	width   int
	height  int
	cells   int
	xyrange float64
	xyscale float64
	zscale  float64
}

func setOptions(values url.Values) *constants {
	var c constants
	c.sin30 = math.Sin(angle)
	c.cos30 = math.Cos(angle)
	c.width = 600
	c.height = 320
	c.cells = 100
	c.xyrange = 30.0
	c.xyscale = float64(c.width) / 2 / c.xyrange
	c.zscale = float64(c.height) * 0.4

	if v, ok := values["width"]; ok {
		num, _ := strconv.Atoi(v[0])
		c.width = num
	}
	if v, ok := values["height"]; ok {
		num, _ := strconv.Atoi(v[0])
		c.height = num
	}
	if v, ok := values["cells"]; ok {
		num, _ := strconv.Atoi(v[0])
		c.cells = num
	}
	if v, ok := values["xyrange"]; ok {
		num, _ := strconv.ParseFloat(v[0], 64)
		c.xyrange = num
	}
	return &c
}

func surface(out io.Writer, c *constants) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", c.width, c.height)
	for i := 0; i < c.cells; i++ {
		for j := 0; j < c.cells; j++ {
			ax, ay := c.corner(i+1, j)
			bx, by := c.corner(i, j)
			cx, cy := c.corner(i, j+1)
			dx, dy := c.corner(i+1, j+1)

			if isSomeInvalid(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func isSomeInvalid(values ...float64) bool {
	for _, v := range values {
		if isInvalid(v) {
			return true
		}
	}
	return false
}

func isInvalid(f float64) bool {
	return math.IsNaN(f) || math.IsInf(f, 0)
}

func (c *constants) corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := c.xyrange * (float64(i)/float64(c.cells) - 0.5)
	y := c.xyrange * (float64(j)/float64(c.cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(c.width)/2 + (x-y)*cos30*c.xyscale
	sy := float64(c.height)/2 + (x+y)*sin30*c.xyscale - z*c.zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	height := math.Sin(r) / r
	return height
}
