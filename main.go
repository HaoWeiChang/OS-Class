package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	mi = 500
	mj = 800
)

var matrixA [mi][mj]float32
var matrixB [mj][mi]float32
var matrixC [mi][mi]float32

func init() {
	start := time.Now()
	initmatrixA(&matrixA)
	initmatrixB(&matrixB)
	duration := time.Since(start).Milliseconds()
	fmt.Printf("初始化陣列計算所花時間%dms\n", duration)
}

func main() {
	start := time.Now()
	forLoop()
	duration := time.Since(start).Milliseconds()
	fmt.Printf("for-loop計算所花時間%dms\n", duration)
	start = time.Now()
	multiplethread1()
	duration = time.Since(start).Milliseconds()
	fmt.Printf("multiplethread1計算所花時間%dms\n", duration)
	start = time.Now()
	multiplethread2()
	duration = time.Since(start).Milliseconds()
	fmt.Printf("multiplethread2計算所花時間%dms\n", duration)
	fmt.Scanln()
}

func forLoop() [mi][mi]float32 {
	for i := 0; i < len(matrixC); i++ {
		for j := 0; j < len(matrixC[0]); j++ {
			matrixC[i][j] = createMatrixC(i, j)
		}
	}
	return matrixC
}

func multiplethread1() [mi][mi]float32 {
	var wg sync.WaitGroup
	for i := 0; i < len(matrixC); i++ {
		for j := 0; j < len(matrixC[0]); j++ {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				matrixC[i][j] = createMatrixC(i, j)
			}(i, j)
		}
	}
	wg.Wait()
	return matrixC
}

func multiplethread2() [mi][mi]float32 {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				istart := i * (len(matrixC) / 5)
				iend := istart + (len(matrixC) / 5)
				jstart := j * (len(matrixC[0]) / 2)
				jend := jstart + (len(matrixC[0]) / 2)
				for i := istart; i < iend; i++ {
					for j := jstart; j < jend; j++ {
						matrixC[i][j] = createMatrixC(i, j)
					}
				}
			}(i, j)
		}
	}
	wg.Wait()
	return matrixC
}

func initmatrixA(mtx *[mi][mj]float32) {
	for i := 1; i <= len(mtx); i++ {
		for j := 1; j <= len(mtx[0]); j++ {
			mtx[i-1][j-1] = 6.6*float32(i) - 3.3*float32(j)
		}
	}
}

func initmatrixB(mtx *[mj][mi]float32) {
	for i := 1; i <= len(mtx); i++ {
		for j := 1; j <= len(mtx[0]); j++ {
			mtx[i-1][j-1] = 100 + 2.2*float32(i) - 5.5*float32(j)
		}
	}
}

func createMatrixC(i int, j int) float32 {
	var res float32 = 0.0
	for x := 0; x < mj; x++ {
		res = res + matrixA[i][x]*matrixB[x][j]
	}
	return res
}
