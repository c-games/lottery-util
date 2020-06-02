package util

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestDecimal(t *testing.T) {

	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromFloat(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	ro := decimal.NewFromFloat(0.12349)

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.088
	fmt.Println("mul:", quantity.Mul(fee))
	//fmt.Println("ro:", ro)

	f, _ := ro.Float64()
	fmt.Printf("%f\n", RoundDown(f, 4))
	fmt.Printf("%f\n", Round(f, 4))

}

// Round returns float64 round number.
//決定捨去或進位
//value
//roundOn 大於等於此數進位，小餘此數捨去
//places 小數點第幾位 從0 開始，0＝第一位
//func Round(val float64, roundOn float64, places int) (newVal float64) {
//	var round float64
//	pow := math.Pow(10, float64(places))
//	digit := pow * val
//	_, div := math.Modf(digit)
//	if div >= roundOn {
//		round = math.Ceil(digit)
//	} else {
//		round = math.Floor(digit)
//	}
//	newVal = round / pow
//	return
//}

func TestFloatMultiply(t *testing.T) {
	type args struct {
		n         float64
		m         float64
		precision int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"0",
			args{0.1111,0.2222,3},
			false,
		},
		{
			"1",
			args{3.10973,99.673,4},
			false,
		},
		{
			"2",
			args{987767.103,199.6998,4},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FloatMultiply(tt.args.n, tt.args.m,tt.args.precision)
			//if (err != nil) != tt.wantErr {
			//	t.Errorf("FloatMultiply() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}

			t.Logf("%f",got)
		})
	}
}
