package calc_path

import (
	"fmt"
)

// временный файл для тестирования ручки, отрисовывает найденный путь просто
// но кажется выводит отзеркаленный))
func parse(path []Transition) {
	// Создаем карту 41x41
	const size = 41
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
		for j := range grid[i] {
			grid[i][j] = '·' // Точка для пустого пространства
		}
	}

	// Отрисовываем путь
	for _, trans := range path {
		// Ставим символ '■' в начальной и конечной точках
		grid[trans.From.Y][trans.From.X] = '■'
		grid[trans.To.Y][trans.To.X] = '■'

		// Рисуем линию между точками
		drawLine(grid, trans.From, trans.To)
	}

	// Выводим карту (переворачиваем Y для корректного отображения)
	for y := size - 1; y >= 0; y-- {
		for x := 0; x < size; x++ {
			fmt.Printf("%c ", grid[y][x])
		}
		fmt.Println()
	}
}

func drawLine(grid [][]rune, From, To Point) {
	// Простая реализация алгоритма Брезенхема для линии
	x0, y0 := int(From.X), int(From.Y)
	x1, y1 := int(To.X), int(To.Y)

	dx := abs(int(x1 - x0))
	dy := abs(int(y1 - y0))
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		grid[y0][x0] = '■'

		if x0 == x1 && y0 == y1 {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
