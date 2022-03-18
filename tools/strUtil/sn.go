package strUtil

import (
	"fmt"
	"github.com/pangu-2/go-tools/tools/datetimeUtil"
	"math"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

//生成单号
//06123xxxxx
//sum 最少10
func MakeYearDaysRand(sum int) string {
	//年
	strs := time.Now().Format("06")
	//一年中的第几天
	days := strconv.Itoa(datetimeUtil.GetDaysInYearByThisYear())
	count := len(days)
	if count < 3 {
		//重复字符0
		days = strings.Repeat("0", 3-count) + days
	}
	//组合
	strs += days
	//剩余随机数
	sum = sum - 5
	if sum < 1 {
		sum = 5
	}
	//0~9999999的随机数
	ran := GetRand()
	pow := math.Pow(10, float64(sum)) - 1
	//fmt.Println("sum=>", sum)
	//fmt.Println("pow=>", pow)
	result := strconv.Itoa(ran.Intn(int(pow)))
	//result := string(ran.Intn(int(math.Pow(0, float64(sum+1)) - 1)))
	count = len(result)
	//fmt.Println("result=>", result)
	if count < sum {
		//重复字符0
		result = strings.Repeat("0", sum-count) + result
	}
	//组合
	strs += result
	return strs
}

//生成订单号
//sum 最少10
func MakeNo(sum int) string {
	return MakeYearDaysRand(sum)
}

//生成订单号
//生成24位订单号
func MakeNo2() string {
	return Generate(time.Now())
}

var num int64

const continuity = "20060102150405"

//BY https://github.com/w3liu/go-common/blob/master/number/ordernum/ordernum.go
//生成24位订单号
//前面17位代表时间精确到毫秒，中间3位代表进程id，最后4位代表序号
func Generate(t time.Time) string {
	s := t.Format(continuity)
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

//对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}
