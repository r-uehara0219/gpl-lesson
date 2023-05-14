package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	dx, dy                 = float64(1.0) / (xmax - xmin), float64(1.0) / (ymax - ymin)
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.
			img.Set(px, py, supersample(x, y))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	// vに二乗と加算を繰り返し、半径2の円の外側に出たら影をつける
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func supersample(x, y float64) color.Color {
	// 4つのサブピクセルの色の平均を返す
	return average(
		average(
			mandelbrot(complex(x-dx, y-dy)),
			mandelbrot(complex(x+dx, y-dy)),
		),
		average(
			mandelbrot(complex(x-dx, y+dy)),
			mandelbrot(complex(x+dx, y+dy)),
		),
	)
}

func average(colors ...color.Color) color.Color {
	var r, g, b, a uint32
	for _, c := range colors {
		_r, _g, _b, _a := c.RGBA()
		r += _r
		g += _g
		b += _b
		a += _a
	}
	n := uint32(len(colors))
	return color.RGBA{uint8(r / n), uint8(g / n), uint8(b / n), uint8(a / n)}
}
