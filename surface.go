package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // Размер канвы в пикселях
	cells         = 100                 // Количество ячеек сетки
	xyrange       = 30.0                // Диапозон осей
	xyscale       = width / 2 / xyrange // Пикселей в еденице x или y
	zscale        = height * 0.4        // Пикселей в еденице z
	angle         = math.Pi / 6         // Углы осей x, y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// Оператор «break» завершает выполнение самого внутреннего оператора «for», «switch» или «select» внутри той же функции.
func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke:grey; fill: white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				break
			}
			bx, by, err := corner(i, j)
			if err != nil {
				break
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				break
			}

			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				break
			}
			fmt.Printf("<polygon points='%g, %g, %g, %g, %g, %g, %g, %g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, error) {
	// Ищем угловую точку (x, y ) ячейки (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхноости z
	z, err := f(x, y)
	if err != nil {
		//log.Fatalf("%s is Infinity Error!!!", err)
		return 0, 0, err
	}
	// Изометрически проецируем (x, y, z) на двумерную канву SVG (sx, sy)
	sx := width/2 + (x+y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) (float64, error) {
	r := math.Hypot(x, y) // Расстояние от (0,0)
	c := math.Sin(r) / r

	// IsInf сообщает, является ли f бесконечностью в соответствии со знаком.
	// Если знак == 0, IsInf сообщает, является ли f бесконечностью. <---
	if math.IsInf(c, 0) == false {
		return math.Round(math.Sin(r) / r), nil
	} else {
		return 0, fmt.Errorf("%s", "Error!!!")
	}
}
