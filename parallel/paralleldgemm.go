package main

import (
	"fmt"
	"sync"
)

const N, T = 4, 4

var w sync.WaitGroup

//))))))interting multiplication laziness for 3

func MatricesMultiplication(N int, B []float64, ch chan []float64) {
	defer w.Done()
	w.Add(1)
	A := <-ch
	C := <-ch

	for i := 0; i < N; i++ { //mult
		C[0] += A[i] * B[i*N]
	}
}

func FillMatrixWithRandomValues(N int, M []float64) {
	for i := 0; i < N; i++ { // col
		for j := 0; j < N; j++ { // row
			M[i*N+j] = float64(i)
		}
	}
}

func OutputMatrix(N int, M []float64) {
	for i := 0; i < N; i++ { // col
		for j := 0; j < N; j++ { // row
			fmt.Printf("%.2f ", M[i*N+j])
		}
		fmt.Println()
	}
}

func main() {

	A, B, C := make([]float64, N*N), make([]float64, N*N), make([]float64, N*N)

	FillMatrixWithRandomValues(N, A)
	FillMatrixWithRandomValues(N, B)

	ch := make(chan []float64)

	for i := 0; i < T; i++ {
		go MatricesMultiplication(N, B, ch)
		ch <- A[i*N : (1+i)*N]
		ch <- C[i*N : (1+i)*N]
	}
	w.Wait()

	OutputMatrix(N, C)
}
