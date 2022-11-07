package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	mi = 50
	mj = 80
)

var matrixA [mi][mj]float32
var matrixB [mj][mi]float32
var forloopMatrix [mi][mi]float32
var multithreadMatrix1 [mi][mi]float32
var multithreadMatrix2 [mi][mi]float32

func init() {
	start := time.Now()
	initMatrixA(&matrixA)
	initMatrixB(&matrixB)
	duration := time.Since(start).Milliseconds()
	fmt.Printf("初始化陣列計算所花時間%dms\n", duration)
}

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("第%d次執行\n", i)
		start := time.Now()
		forLoop(&forloopMatrix)
		duration := time.Since(start).Milliseconds()
		fmt.Printf("for-loop計算所花時間%dms\n", duration)

		start = time.Now()
		multiplethread1(&multithreadMatrix1)
		duration = time.Since(start).Milliseconds()
		fmt.Printf("multiplethread1計算所花時間%dms\n", duration)

		start = time.Now()
		multiplethread2(&multithreadMatrix2)
		duration = time.Since(start).Milliseconds()
		fmt.Printf("multiplethread2計算所花時間%dms\n\n", duration)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Print("請按下任意鍵離開")
	fmt.Scanln()
}

func forLoop(mtx *[mi][mi]float32) {
	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[0]); j++ {
			mtx[i][j] = createMatrixC(i, j)
		}
	}
}

func multiplethread1(mtx *[mi][mi]float32) {
	var wg sync.WaitGroup
	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[0])/50; j++ {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				jstart := j * 50
				jend := jstart + 50
				for j := jstart; j < jend; j++ {
					mtx[i][j] = createMatrixC(i, j)
				}
			}(i, j)
		}
	}
	wg.Wait()
}

func multiplethread2(mtx *[mi][mi]float32) {
	var wg sync.WaitGroup
	for i := 0; i < len(mtx)/10; i++ {
		for j := 0; j < len(mtx)/25; j++ {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				istart := i * 10
				iend := istart + 10
				jstart := j * 25
				jend := jstart + 25
				for i := istart; i < iend; i++ {
					for j := jstart; j < jend; j++ {
						mtx[i][j] = createMatrixC(i, j)
					}
				}
			}(i, j)
		}
	}
	wg.Wait()
}

func initMatrixA(mtx *[mi][mj]float32) {
	for i := 1; i <= len(mtx); i++ {
		for j := 1; j <= len(mtx[0]); j++ {
			mtx[i-1][j-1] = 6.6*float32(i) - 3.3*float32(j)
		}
	}
}

func initMatrixB(mtx *[mj][mi]float32) {
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
