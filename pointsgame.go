package main

import (
	"fmt"
	"math"
)

var Threshold float64 = 1e-6
var CardNumber int = 4

var ResultValue int = 24

var result [4]string

var number [4]float64

func output(a interface{}) {
	fmt.Println("print----", a)
}

func PointsGame(n int) bool {
	if n == 1 {
		if math.Abs(number[0]-float64(ResultValue)) < Threshold {
			output(result[0])
			return true

		}

		return false
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			var a, b float64
			var expa, expb string
			a = number[i]
			b = number[j]
			number[j] = number[n-1]

			expa = result[i]
			expb = result[j]
			result[j] = result[n-1]
			result[i] = fmt.Sprintf("(%s+%s)", expa, expb)

			number[i] = a + b
			if PointsGame(n - 1) {
				return true
			}
			result[i] = fmt.Sprintf("(%s-%s)", expa, expb)
			number[i] = a - b
			if PointsGame(n - 1) {
				return true
			}

			result[i] = fmt.Sprintf("(%s-%s)", expb, expb)
			number[i] = b - a
			if PointsGame(n - 1) {
				return true
			}

			result[i] = fmt.Sprintf("(%s*%s)", expa, expb)
			number[i] = a * b
			if PointsGame(n - 1) {
				return true
			}
			if b != 0 {
				result[i] = fmt.Sprintf("(%s/%s)", expa, expb)
				number[i] = a / b
				if PointsGame(n - 1) {
					return true
				}
			}
			if a != 0 {
				result[i] = fmt.Sprintf("(%s/%s)", expb, expa)
				number[i] = b / a
				if PointsGame(n - 1) {
					return true
				}
			}

			number[i] = a
			number[j] = b
			result[i] = expa
			result[j] = expb

		}

	}
	return false
}

func main() {
	// var x int
	number = [4]float64{1, 8, 3, 5}
	result = [4]string{"1", "8", "3", "5"}
	if PointsGame(4) {
		output("success")

	} else {
		output("fail")
	}

}
