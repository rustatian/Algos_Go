package main

import (
	"math"
)

func main() {

}

func F(x float64) float64 {
	return x / (1.0 + math.Sqrt(1.0 + x))
}
