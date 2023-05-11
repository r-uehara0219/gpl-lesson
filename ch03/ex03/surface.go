package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	topColor      = "#ff0000"
	bottomColor   = "#0000ff"

	// 実際の頂点の座標と異なるが、図示したときの見た目を考慮して広く設定
	maxHeight = 0.9
	minHeight = -0.2
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			if isSomeInvalid(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color(az, bz, cz, dz))
		}
	}
	fmt.Println("</svg>")
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

// z座標を戻り値に追加
func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	height := math.Sin(r) / r
	return height
}

func color(az, bz, cz, dz float64) string {
	height := (az + bz + cz + dz) / 4
	if height > maxHeight {
		return topColor
	} else if height < minHeight {
		return bottomColor
	} else {
		return "#000000"
	}
}
