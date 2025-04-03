package main

import "math/rand/v2"

// Represents a matrix.
type Matrix struct {
	rows int
	cols int
	data [][]float64
}

// Creates a new matrix with the given number of rows and columns.
func NewMatrix(rows int, cols int) (*Matrix, error) {
	mat := &Matrix{}
	mat.rows = rows
	mat.cols = cols
	mat.data = make([][]float64, rows)

	// Populate the matrix with random floats from 0.0 to 1.0
	for i := range mat.data {
		mat.data[i] = make([]float64, cols)
		for j := range mat.data[i] {
			mat.data[i][j] = rand.Float64()
		}
	}
	return mat, nil
}

// Returns the number of rows in the matrix.
func (m *Matrix) Rows() int {
	return m.rows
}

// Returns the number of columns in the matrix.
func (m *Matrix) Cols() int {
	return m.cols
}
