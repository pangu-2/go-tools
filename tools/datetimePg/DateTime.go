package datetimePg

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

//https://www.jianshu.com/p/7f691bbc9b85
// 数据库日期时间映射
type DateTime time.Time

//MarshalJSON 实现了 Marshaler 接口
func (t DateTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", t.DateTime())
	return []byte(stamp), nil
}

//UnmarshalJSON 实现了 Unmarshaler 接口
func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = DateTime(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+Y_M_D_H_I_S+`"`, string(data))
	*t = DateTime(now)
	return
}

//日期时间
func (t DateTime) DateTime() string {
	return t.Format(Y_M_D_H_I_S)
}

//日期
func (t DateTime) Date() string {
	return t.Format(Y_M_D)
}

//时间
func (t DateTime) Time() string {
	return t.Format(H_I_S)
}
func (t DateTime) Format(layout string) string {
	return Format(t, layout)
}

// 用于 fmt.Println 和后续验证场景
func (t DateTime) String() string {
	return time.Time(t).Format(Y_M_D_H_I_S)
}

// 写入 mysql 时调用
func (t DateTime) Value() (driver.Value, error) {
	if &t == nil {
		return nil, nil
	}
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(Y_M_D_H_I_S)), nil
}

// 检出 mysql 时调用
func (t *DateTime) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	//https://blog.csdn.net/bailaoshi666/article/details/114921633
	switch v.(type) {
	case time.Time:
		tTime, _ := time.Parse(Y_M_D_H_I_S_CST, v.(time.Time).String())
		*t = DateTime(tTime)
	case DateTime:
		*t = v.(DateTime)
	case *DateTime:
		v1 := v.(*DateTime)
		*t = *v1
	default:
		return errors.New("can not convert to Date")
	}
	return nil
}
