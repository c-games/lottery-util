package util

import (
	"github.com/shopspring/decimal"
)

//FloatMultiply 計算兩個福點數乘法，
//取精確度precision(小樹後幾位	)
//precision 精確度。保留小數點後幾位
func FloatMultiply(n, m float64, precision int32) float64 {
	nn := decimal.NewFromFloat(n)
	mm := decimal.NewFromFloat(m)
	nm := nn.Mul(mm)

	n4 := nm.Truncate(precision)

	f64, _ := n4.Float64()
	return f64
}
