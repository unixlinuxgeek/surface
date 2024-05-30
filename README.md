### Упражнение 3.2

Поэксперементируйте с визуализациями других функции из пакета math.
Сможете ли вы получить изображения наподобие коробки для яйц, седла или холма?

Используем перенаправление вывода: 
```
$ go run ./surface.go 1 > 1.svg
$ go run ./surface.go 2 > 2.svg
$ go run ./surface.go 3 > 3.svg
$ go run ./surface.go 4 > 4.svg
$ go run ./surface.go 5 > 5.svg
$ go run ./surface.go 6 > 6.svg
$ go run ./surface.go 7 > 7.svg
$ go run ./surface.go 8 > 8.svg
$ go run ./surface.go 9 > 9.svg
```

```
$ go run ./surface.go 10
Не введены аргументы (от 1 до 9)!!!
```