package strPg

import (
	"fmt"
	"github.com/pangu-2/go-tools/tools/convPg"
	"github.com/pangu-2/go-tools/tools/datetimePg"
	"github.com/pangu-2/go-tools/tools/slicePg"
	"math"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

// 生成单号
// 06123xxxxx
// sum 最少10
func MakeYearDaysRand(sum int) string {
	//年
	strs := time.Now().Format("06")
	//一年中的第几天
	days := strconv.Itoa(datetimePg.GetDaysInYearByThisYear())
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

// 生成订单号
// sum 最少10
func MakeNo(sum int) string {
	return MakeYearDaysRand(sum)
}

// 生成订单号
// 生成24位订单号
func MakeNo2() string {
	return Generate(time.Now())
}

var num int64

const continuity = "20060102150405"

// BY https://github.com/w3liu/go-common/blob/master/number/ordernum/ordernum.go
// 生成24位订单号
// 前面17位代表时间精确到毫秒，中间3位代表进程id，最后4位代表序号
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

// 对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

// GenerateNumberId22 生成22位时间格式id号，例子:2303050010557521234567
// 格式：年(2位)+月+日+时+分+秒+毫秒(3位)+随机数字(7位)
func GenerateNumberId22() string {
	now := time.Now()
	format := now.Format(datetimePg.YMDHIS_SSS2)
	format = strings.Replace(format, ".", "", -1)
	//合并
	format = format + GetNanoIdNumber(7)
	//转换
	return format
}

// GenerateNumberId18 生成22位时间格式id号，例子:2303050010557521234567
// 格式：年(2位)+月+日+时+分+秒+毫秒(3位)+随机数字(3位)
func GenerateNumberId18() int64 {
	now := time.Now()
	format := now.Format(datetimePg.YMDHIS_SSS2)
	format = strings.Replace(format, ".", "", -1)
	//合并
	format = format + GetNanoIdNumber(3)
	//转换
	return convPg.StrToInt64(format)
}

// GenerateNumberId 生成指定时间的n位时间格式id号，例子:2303050010557521234567
// 格式：年(2位)+月+日+时+分+秒+毫秒(3位)+随机数字(7位)
// 最低位数位 15位
func GenerateNumberId(now time.Time, len int) string {
	if len < 15 {
		len = 15
	}
	format := now.Format(datetimePg.YMDHIS_SSS2)
	format = strings.Replace(format, ".", "", -1)
	//合并
	if len > 15 {
		format = format + GetNanoIdNumber(len-15)
	}
	//转换
	return format
}

func GenerateUnixId(now time.Time, len int) string {
	if len < 13 {
		len = 13
	}
	milli := now.UnixMilli()
	formatInt := strconv.FormatInt(milli, 10)
	//合并
	if len > 13 {
		return formatInt + GetNanoIdNumber(len-13)
	}
	//转换
	return formatInt
}

// GenerateUnixIdIntString  生成linux时间戳格式id号，例子:167798942647399999
// 格式：当前linux时间戳(10位秒数)+毫秒(3位)+随机数值(5位)
func GenerateUnixIdIntString() int64 {
	//13位 毫秒
	milli := time.Now().UnixMilli()
	formatInt := strconv.FormatInt(milli, 10)
	//转换
	return convPg.StrToInt64(formatInt + GetNanoIdNumber(5))
}

// GenerateUnixIdInt  生成linux时间戳格式id号，例子:167798942647399999
// 格式：当前linux时间戳(10位秒数)+毫秒(3位)+随机数值(5位)
func GenerateUnixIdInt() int64 {
	//13位 毫秒
	milli := time.Now().UnixMilli()
	formatInt := strconv.FormatInt(milli, 10)
	//转换
	return convPg.StrToInt64(formatInt + GetNanoIdNumber(5))
}

// GenerateUnixIdIntMore  生成linux时间戳格式id号，例子:167798942647399999
// 格式：当前linux时间戳(10位秒数)+毫秒(3位)+随机数值(5位)
func GenerateUnixIdIntMore(count int) []int64 {
	//13位 毫秒
	milli := time.Now().UnixMilli()
	formatInt := strconv.FormatInt(milli, 10)
	slice := make([]int64, 0)
	mapT := make(map[string]bool)
	for i := 1; i <= count; i++ {
		idInt := GetNanoIdNumber(5)
		if _, ok := mapT[idInt]; ok {
			//重复自动 新生成
			count++
			continue
		} else {
			mapT[idInt] = true
		}
		slice = append(slice, convPg.StrToInt64(formatInt+idInt))
	}
	return slice
}

// GenerateUnixIdIntMoreUnique 生成linux时间戳格式id号，例子:167798942647399999
func GenerateUnixIdIntMoreUnique(sum int) []int64 {
	idArr := make([]int64, 0)
	idArr = GenerateUnixIdIntMore(sum)
	//去除重复的
	idArr = slicePg.Unique(idArr)
	sumUnique := len(idArr)
	if sum > sumUnique {
		more2 := GenerateUnixIdIntMore(sum - sumUnique)
		idArr = slicePg.Merge(idArr, more2)
	}
	return idArr
}
