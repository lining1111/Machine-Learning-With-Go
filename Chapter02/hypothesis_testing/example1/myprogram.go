package main

import (
	"fmt"

	"github.com/gonum/stat"
)

func main() {

	// Define observed and expected values. Most
	// of the time these will come from your
	// data (website visits, etc.).
	observed := []float64{48, 52}
	expected := []float64{50, 50}

	// Calculate the ChiSquare test statistic.
	//卡方统计量，越接近0,说明观察值和期望值越接近
	//自由度为k的p值，表示大到什么程度才拒绝观察值
	chiSquare := stat.ChiSquare(observed, expected)

	fmt.Println(chiSquare)
}
