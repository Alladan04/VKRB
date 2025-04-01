package model_manager

import "mivar_robot_api/internal/entity"

// GetExits возвращает координаты всех граничных элементов матрицы, равных 0 (потенциальных выходов из лабиринта).
func (m *Manager) GetExits(matrix [][]uint8) []entity.Point {
	var zeros []entity.Point

	if len(matrix) == 0 {
		return zeros
	}

	rows := len(matrix)
	cols := len(matrix[0])

	// Проверяем верхнюю и нижнюю границы
	for col := 0; col < cols; col++ {
		// Верхняя строка
		if matrix[0][col] == 0 {
			zeros = append(zeros, entity.Point{X: int64(0), Y: int64(col)})
		}
		// Нижняя строка (если есть)
		if rows > 1 && matrix[rows-1][col] == 0 {
			zeros = append(zeros, entity.Point{X: int64(rows - 1), Y: int64(col)})
		}
	}

	// Проверяем левую и правую границы (исключая углы, чтобы не дублировать)
	for row := 1; row < rows-1; row++ {
		// Левый столбец
		if matrix[row][0] == 0 {
			zeros = append(zeros, entity.Point{X: int64(row), Y: int64(0)})
		}
		// Правый столбец (если есть)
		if cols > 1 && matrix[row][cols-1] == 0 {
			zeros = append(zeros, entity.Point{X: int64(row), Y: int64(cols - 1)})
		}
	}

	return zeros
}

// GetExitsByModelID возвращает координаты всех граничных элементов матрицы, равных 0 (потенциальных выходов из лабиринта).
func (m *Manager) GetExitsByModelID(modelID string) ([]entity.Point, error) {
	matrix, err := m.inMemRepo.GetLabirintFromCache(modelID)
	if err != nil {
		return nil, err
	}

	return m.GetExits(matrix), nil
}
