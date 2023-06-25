package main

import "math"
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for y := range s {
		s[y] = make([]uint8, dx)
		for x := range s[y] {
			s[y][x] = gopher(x, y)
		}
	}

	return s
}

func gopher(x, y int) uint8 {
    xx := float64(x)
	yy := float64(y)
	
	params := [][6]float64 {
		{180, 50, 3, 3, 0, 255},
		{175, 50, 10, 10, 0, 0},
		{105, 80, 3, 3, 0, 255},
		{100, 80, 10, 10, 0, 0},
		{155, 55, 30, 30, 0, 255},
		{155, 55, 32, 32, 0, 0},
		{80, 85, 30, 30, 0, 255},
		{80, 85, 32, 32, 0, 0},
		{130, 90, 7, 1, 24, 255},
		{125, 95, 13, 8, 24, 0},
		{133, 128, 12, 5, 114, 255},
		{133, 128, 13, 6, 114, 0},
		{143, 124, 12, 5, 116, 255},
		{143, 124, 13, 6, 116, 0},
		{130, 108, 22, 12, 23, 160},
		{130, 108, 24, 14, 23, 0},
		{183, 178, 14, 8, -20, 255},
		{183, 178, 15, 9, -20, 0},
		{131, 108, 96, 100, 0, 220},
		{151, 158, 90, 100, 0, 220},
		{151, 188, 90, 100, 0, 220},
		{243, 148, 14, 8, -20, 255},
		{243, 148, 15, 9, -20, 0},
		{205, 36, 2, 2, 0, 255},
		{203, 42, 10, 10, 0, 0},
		{206, 44, 21, 21, 0, 255},
		{206, 44, 23, 23, 0, 0},
		{33, 102, 10, 10, 0, 0},
		{36, 104, 21, 21, 0, 255},
		{36, 104, 23, 23, 0, 0},
	}
    
	for _, p := range params {
		if judge(xx, yy, p[0], p[1], p[2], p[3], p[4]) {
			return uint8(p[5])
		}
	}
	
	return 255
}

func judge(x, y, cx, cy, a, b, deg float64) bool {
	t := deg * math.Pi / 180
	ct := math.Cos(t)
	st := math.Sin(t)
	x2 := (x - cx)
	y2 := (y - cy)
	x3 := (ct * x2 - st * y2) / a
	y3 := (st * x2 + ct * y2) / b
    return x3 * x3 + y3 * y3 < 1
}

func main() {
	pic.Show(Pic)
}