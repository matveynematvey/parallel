package main

import (
	"fmt"
	"math/rand"
)

func MatricesMultiplication(M, N, K int, A, B, C *[]float64) {

	for i := 0; i < N; i++ { //col
		for j := 0; j < M; j++ { //row
			for k := 0; k < K; k++ { //mult
				(*C)[i*M+j] += (*A)[k*M+j] * (*B)[i*K+k]
			}
		}
	}

}

func FillMatrixWithRandomValues(C, R int, M *[]float64) {
	for i := 0; i < C; i++ { // col
		for j := 0; j < R; j++ { // row
			(*M)[i*R+j] = rand.Float64() * 10.0
		}
	}
}

func OutputMatrix(C, R int, M *[]float64) {
	for i := 0; i < C; i++ { // col
		for j := 0; j < R; j++ { // row
			fmt.Printf("%.2f ", (*M)[i*R+j])
		}
		fmt.Println()
	}
}

func main() {
	const M, K, N int = 4, 2, 3 //matrices size

	A, B, C := make([]float64, M*K), make([]float64, K*N), make([]float64, M*N)

	FillMatrixWithRandomValues(K, M, &A)
	FillMatrixWithRandomValues(N, K, &B)

	MatricesMultiplication(M, N, K, &A, &B, &C)

	OutputMatrix(N, M, &C)
}
