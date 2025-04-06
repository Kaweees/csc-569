package main

import (
	"fmt"
	"testing"
)

// BenchmarkMatrixMultiply benchmarks the matrix multiplication function
// with matrices of different sizes
func BenchmarkMatrixMultiplyNaive(b *testing.B) {
	// Run benchmarks with different matrix sizes
	benchmarkSizes := []int{10, 50, 100, 200, 500, 1000}

	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("%dx%d", size, size), func(b *testing.B) {
			// Create two square matrices of the given size
			m1, _ := NewMatrix(size, size)
			m2, _ := NewMatrix(size, size)

			// Reset the timer before the actual benchmark loop
			b.ResetTimer()

			// Run the multiplication b.N times
			for i := 0; i < b.N; i++ {
				m1.MultiplyNaive(m2)
			}
		})
	}
}

func BenchmarkMatrixMultiplyRoutines(b *testing.B) {
	// Run benchmarks with different matrix sizes
	benchmarkSizes := []int{10, 50, 100, 200, 500, 1000}

	for _, size := range benchmarkSizes {
		b.Run(fmt.Sprintf("%dx%d", size, size), func(b *testing.B) {
			// Create two square matrices of the given size
			m1, _ := NewMatrix(size, size)
			m2, _ := NewMatrix(size, size)

			// Reset the timer before the actual benchmark loop
			b.ResetTimer()

			// Run the multiplication b.N times
			for i := 0; i < b.N; i++ {
				m1.MultiplyRoutines(m2)
			}
		})
	}
}

// // BenchmarkMatrixMultiplyNonSquare benchmarks multiplication with non-square matrices
// func BenchmarkMatrixMultiplyNonSquare(b *testing.B) {
// 	// Test different matrix size combinations
// 	sizeCombinations := []struct {
// 		rows1 int
// 		cols1 int
// 		cols2 int
// 	}{
// 		{100, 50, 200},  // 100x50 * 50x200
// 		{500, 100, 300}, // 500x100 * 100x300
// 		{1000, 20, 50},  // 1000x20 * 20x50
// 	}

// 	for _, sc := range sizeCombinations {
// 		name := fmt.Sprintf("%dx%d*%dx%d", sc.rows1, sc.cols1, sc.cols1, sc.cols2)
// 		b.Run(name, func(b *testing.B) {
// 			m1, _ := NewMatrix(sc.rows1, sc.cols1)
// 			m2, _ := NewMatrix(sc.cols1, sc.cols2)

// 			b.ResetTimer()

// 			for i := 0; i < b.N; i++ {
// 				m1.MultiplyNaive(m2)
// 			}
// 		})
// 	}
// }
