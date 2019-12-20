// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)


func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
    epsX = (xmax - xmin)/width
    epsY = (ymax - ymin)/height
	)
  
  offsetX := []float64{-epsX,epsX}
  offsetY := []float64{-epsY,epsY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			//z := complex(x, y)
      // One dot to 4 dot and get 4 mandelbrot color and 
      // then get avg color
      dotSet := make([]color.Color,0)
      for i := 0; i<2; i++ {
        for j :=0; j<2; j++ {
           dotY := y + offsetY[i]
           dotX := x + offsetX[j]
           dotZ := complex(dotX, dotY)
           dotSet = append(dotSet,mandelbrot(dotZ)) 
        }
      }

			// Image point (px, py) represents complex value z.
			//img.Set(px, py, mandelbrot(z))
			img.Set(px, py, avgMandelbrot(dotSet))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast * n,255 - contrast*n,255,255}
			//return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func avgMandelbrot(dotSet []color.Color) color.Color {
  var r,g,b,a uint8
  for _,v := range dotSet {
    _r,_g,_b,_a := v.RGBA()
    r += uint8(_r)
    g += uint8(_g)
    b += uint8(_b)
    a += uint8(_a)
  }
  lenDot := uint8(len(dotSet))
  return color.RGBA{r/lenDot, g/lenDot, b/lenDot, a/lenDot}
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
