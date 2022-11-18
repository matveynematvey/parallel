package main

import (
	"fmt"
	"sync"
)

const N int = 32

func main() {
	arr1, arr2, res := make([]int, N, N), make([]int, N, N), make([]int, N, N)

	for ind := 0; ind < N; ind++ {
		arr1[ind], arr2[ind] = ind, ind
	}

	var wait sync.WaitGroup
	for i := 0; i < 4; i++ {
		wait.Add(1)
		start, end := i*N/4, (i+1)*N/4
		go func(arr1, arr2, res []int) {
			defer wait.Done()
			for ind := range arr1 {
				res[ind] = arr1[ind] + arr2[ind]
			}
		}(arr1[start:end], arr2[start:end], res[start:end])
	}
	wait.Wait()
	fmt.Println(res)
}
