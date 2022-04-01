package convPg

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestSumFloat64List(t *testing.T) {
	var randList []float64 = []float64{123123.1, 14123.5}
	var amount int64 = 2342
	var sum float64
	decimalSum := decimal.NewFromFloat(sum)
	for _, item := range randList {
		decimalItem := decimal.NewFromFloat(item)
		decimalSum = decimalSum.Add(decimalItem)
	}
	sum, _ = decimalSum.Float64()
	if sum != float64(amount) {
		fmt.Println(sum)
		t.Error("结果不正确")
	}
	//3.1 + 2 float和int相加
	var n1 float64 = 3.1
	var n2 int = 2
	d1 := decimal.NewFromFloat(n1).Add(decimal.NewFromFloat(float64(n2)))
	fmt.Println(d1)
	//
	test := decimal.NewFromFloat(16.25).Div(decimal.NewFromFloat(float64(8)))
	fmt.Println(test)
	// Round() 可以进行四舍五入
	s := decimal.NewFromFloat(5.45).Round(1).String() // output: "5.5"
	fmt.Println(s)

	XX, err := decimal.NewFromString("121313131312312.1231234234232453242")
	fmt.Println(XX)
	fmt.Println(err)
}
