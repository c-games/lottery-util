package util

import (
	"math"
	"math/rand"
	"time"
)

//RandomInt return ran int
//include max
func RandomInt(min, max int) int {
	if min >= max {
		return min
	}

	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min+1)
}

//4捨5入
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

//直接捨去
func RoundDown(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc(f*pow10_n) / pow10_n
}