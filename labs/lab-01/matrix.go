package main

import (
	"fmt"
	"math/rand"
)

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

// Returns the element at the given row and column.
func (m *Matrix) Get(row int, col int) float64 {
	return m.data[row][col]
}

// Sets the element at the given row and column.
func (m *Matrix) Set(row int, col int, value float64) {
	m.data[row][col] = value
}

// Returns the matrix as a string.
func (m *Matrix) String() string {
	return fmt.Sprintf("%v", m.data)
}

func (m *Matrix) MultiplyNaive(other *Matrix) (*Matrix, error) {
	if m.cols != other.rows {
		return nil, fmt.Errorf("matrix dimensions incompatible for multiplication: first matrix is %dx%d, second is %dx%d", m.rows, m.cols, other.rows, other.cols)
	}
	res, err := NewMatrix(m.rows, other.cols)
	if err != nil {
		return nil, fmt.Errorf("failed to create result matrix: %v", err)
	}

	// Perform matrix multiplication
	for i := 0; i < m.rows; i++ {
		for j := 0; j < other.cols; j++ {
			for k := 0; k < m.cols; k++ {
				res.data[i][j] += m.data[i][k] * other.data[k][j]
			}
		}
	}
	return res, nil
}

// Perform a matrix multiplication using goroutines
func (m *Matrix) MultiplyRoutines(other *Matrix) (*Matrix, error) {
	if m.cols != other.rows {
		return nil, fmt.Errorf("matrix dimensions incompatible for multiplication: first matrix is %dx%d, second is %dx%d", m.rows, m.cols, other.rows, other.cols)
	}
	res, err := NewMatrix(m.rows, other.cols)
	if err != nil {
		return nil, fmt.Errorf("failed to create result matrix: %v", err)
	}

	// Create a channel to wait for all goroutines to complete
	done := make(chan bool)

	// For each row in the result matrix, spawn a goroutine
	for i := 0; i < m.rows; i++ {
		go func(row int) {
			// Calculate each element in this row
			for j := 0; j < other.cols; j++ {
				sum := 0.0
				for k := 0; k < m.cols; k++ {
					sum += m.data[row][k] * other.data[k][j]
				}
				res.data[row][j] = sum
			}
			// Signal completion
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < m.rows; i++ {
		<-done
	}

	return res, nil
}
