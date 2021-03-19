package matrix

func Erase(matrix [][]int) [][]int {
	rows := make([]bool, len(matrix))
	columns := make([]bool, len(matrix[0]))
	for i, ints := range matrix {
		for j, elem := range ints {
			if elem == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}

	for i, row := range rows {
		if row {
			matrix = EraseRow(matrix, i)
		}
	}

	for i, column := range columns {
		if column {
			matrix = EraseColumn(matrix, i)
		}
	}

	return matrix
}

func EraseColumn(matrix [][]int, column int) [][]int {
	for _, ints := range matrix {
		ints[column] = 0
	}
	return matrix
}

func EraseRow(matrix [][]int, row int) [][]int {
	matrix[row] = make([]int, len(matrix[row]))
	return matrix
}
