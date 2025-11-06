package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	strs := strings.Split(s, "\n")
	matrix := make(Matrix, 0, len(strs))
	for _, str := range strs {
		numStrs := strings.Fields(str)
		lenRow := len(numStrs)
		if lenRow == 0 {
			return nil, fmt.Errorf("empty row")
		}
		if len(matrix) > 0 && lenRow != len(matrix[0]) {
			return nil, fmt.Errorf("inconsistent row lengths")
		}
		row := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", numStr)
			}
			row[i] = num
		}
		matrix = append(matrix, row)
	}
	return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
		return [][]int{}
	}
	numCols := len(m[0])
	cols := make([][]int, numCols)
	for c := 0; c < numCols; c++ {
		col := make([]int, len(m))
		for r := 0; r < len(m); r++ {
			col[r] = m[r][c]
		}
		cols[c] = col
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i := range m {
		rows[i] = make([]int, len(m[i]))
		copy(rows[i], m[i])
	}
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
