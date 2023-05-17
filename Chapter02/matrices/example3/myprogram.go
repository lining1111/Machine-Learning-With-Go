package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
)

func main() {

	// Create a new matrix a. 这是一个R矩阵
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	fa := mat.Formatted(a, mat.Prefix("      "))
	fmt.Printf("a = %v\n\n", fa)
	// Compute and output the transpose of the matrix.
	ft := mat.Formatted(a.T(), mat.Prefix("      "))
	fmt.Printf("a^T = %v\n\n", ft)

	// Compute and output the determinant of a. 矩阵的行列式，决定了矩阵的解是否唯一
	//这里url：https://www.cnblogs.com/tsingke/p/10671318.html说的就很好
	deta := mat.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	// Compute and output the inverse of a.
	aInverse := mat.NewDense(0, 0, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat.Formatted(aInverse, mat.Prefix("       "))
	fmt.Printf("a^-1 = %v\n\n", fi)
	b := mat.NewDense(0, 0, nil)
	b.Mul(a, aInverse)
	fb := mat.Formatted(b, mat.Prefix("       "))
	fmt.Printf("b = %v\n\n", fb)
}
