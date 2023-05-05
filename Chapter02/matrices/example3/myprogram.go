package main

import (
	"fmt"
	"log"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// Create a new matrix a. 这是一个R矩阵
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	fa := mat64.Formatted(a, mat64.Prefix("      "))
	fmt.Printf("a = %v\n\n", fa)
	// Compute and output the transpose of the matrix.
	ft := mat64.Formatted(a.T(), mat64.Prefix("      "))
	fmt.Printf("a^T = %v\n\n", ft)

	// Compute and output the determinant of a. 矩阵的行列式，决定了矩阵的解是否唯一
	//这里url：https://www.cnblogs.com/tsingke/p/10671318.html说的就很好
	deta := mat64.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	// Compute and output the inverse of a.
	aInverse := mat64.NewDense(0, 0, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat64.Formatted(aInverse, mat64.Prefix("       "))
	fmt.Printf("a^-1 = %v\n\n", fi)
	b := mat64.NewDense(0, 0, nil)
	b.Mul(a, aInverse)
	fb := mat64.Formatted(b, mat64.Prefix("       "))
	fmt.Printf("b = %v\n\n", fb)
}
