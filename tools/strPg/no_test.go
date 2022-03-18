package strPg

import "testing"

func Test_No(t *testing.T) {
	str := MakeYearDaysRand(20)
	t.Log(str)
}
