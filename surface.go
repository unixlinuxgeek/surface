// Упражнение 3.2
//
// Поэксперементируйте с визуализациями других функции из пакета math.
// Сможете ли вы получить изображения наподобие коробки для яйц, седла или холма?

package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	width, height = 600, 320            // Размер канвы в пикселях
	cells         = 100                 // Количество ячеек сетки
	xyrange       = 30.0                // Диапозон осей
	xyscale       = width / 2 / xyrange // Пикселей в еденице genSVG или y
	zscale        = height * 0.4        // Пикселей в еденице z
	angle         = math.Pi / 6         // Углы осей genSVG, y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// Оператор «break» завершает выполнение самого внутреннего оператора «for», «switch» или «select» внутри той же функции.
func main() {
	fmt.Println(len(os.Args))
	if len(os.Args) > 1 {
		n, _ := strconv.Atoi(os.Args[1:2][0])
		genSVG(n)
	} else {
		log.Fatal("Введите аргументы (от 1 до 9) !!!\n")
	}
}

func genSVG(n int) {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke:grey; fill: white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, n)
			bx, by := corner(i, j, n)
			cx, cy := corner(i, j+1, n)
			dx, dy := corner(i+1, j+1, n)

			fmt.Printf("<polygon points='%g, %g, %g, %g, %g, %g, %g, %g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j, op int) (float64, float64) {
	// Ищем угловую точку (genSVG, y ) ячейки (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Вычисляем высоту поверхноости z
	var z float64
	switch op {
	case 1:
		z = f1(x, y)
	case 2:
		z = f2(x, y)
	case 3:
		z = f3(x, y)
	case 4:
		z = f4(x, y)
	case 5:
		z = f5(x, y)
	case 6:
		z = f6(x, y)
	case 7:
		z = f7(x, y)
	case 8:
		z = f8(x, y)
	case 9:
		z = f9(x, y)
	default:
		log.Fatal("Need adding correct argument between 1-9 !!!")
	}

	// Изометрически проецируем (genSVG, y, z) на двумерную канву SVG (sx, sy)
	sx := width/2 + (x+y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// Hypot возвращает Sqrt(p*p + q*q), стараясь избежать ненужного переполнения или потери значения.
func f1(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от (0,0)
	return math.Sin(r) / r
}

// math.Dim возвращает максимум genSVG-y или 0.
func f2(x, y float64) float64 {
	r := math.Dim(x, y)
	return math.Sin(r) / r
}

// math.Atan2 возвращает арктангенс y/genSVG, используя знаки двойки для определения квадранта возвращаемого значения.
func f3(x, y float64) float64 {
	r := math.Atan2(y, x)
	return math.Sin(r) / r
}

// math.Max возвращает большее из genSVG или y.
func f4(x, y float64) float64 {
	r := math.Max(y, x)
	return math.Sin(r) / r
}

// math.Nextafter возвращает следующее представимое значение float64 после x в сторону y.
func f5(x, y float64) float64 {
	r := math.Nextafter(y, x)
	return math.Sin(r) / r
}

// math.Remainder возвращает остаток с плавающей запятой IEEE 754 от x/y.
func f6(x, y float64) float64 {
	r := math.Remainder(y, x)
	return math.Sin(r) / r
}

// math.Pow возвращает x**y, экспоненту по основанию x от y.
func f7(x, y float64) float64 {
	r := math.Pow(y, x)
	return math.Sin(r) / r
}

// math.Min возвращает меньшее из x или y.
func f8(x, y float64) float64 {
	r := math.Min(y, x)
	return math.Sin(r) / r
}

// math.Mod возвращает остаток с плавающей запятой от genSVG/y. Величина результата меньше y, а его знак соответствует знаку genSVG.
func f9(x, y float64) float64 {
	r := math.Mod(y, x)
	return math.Sin(r) / r
}
