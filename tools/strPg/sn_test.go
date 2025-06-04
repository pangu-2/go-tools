package strPg

import (
	"fmt"
	"github.com/pangu-2/go-tools/tools/convPg"
	"strconv"
	"testing"
)

func TestSn(t *testing.T) {
	id := "167798894487899999"
	fmt.Println(len(id))
	fmt.Println(len("18446744073709551615"))
	fmt.Println(convPg.StrToInt64(id))
	fmt.Println(strconv.ParseUint(id, 10, 64))
	fmt.Println(GenerateNumberId22())
	fmt.Println(GenerateUnixIdInt())
	mapT := make(map[int64]int64)
	count := 0
	for i := 1; i <= 1000; i++ {
		idInt := GenerateUnixIdInt()
		if item, ok := mapT[idInt]; ok {
			fmt.Println("重复", item)
			count++
		} else {
			mapT[idInt] = idInt
		}
	}
	fmt.Println(mapT)
	fmt.Println(count)
	fmt.Println(GenerateUnixIdIntMore(10000))
}

func TestSn2(t *testing.T) {
	fmt.Println(MakeNo(20))
	//25 155 340161347063982
	fmt.Println(GenerateNumberId22())
	//2506041947215239997347
	//9223372036854775808
	//18446744073709551615   uint64 最大值
	fmt.Println(GenerateNumberId19())
	//2506041952171494613
}
