package datetimePg

import "time"

//Unix时间戳
func UnixTime() int64 {
	return time.Now().Unix()
}

func UnixTimeFormat(val int64, str string) string {
	return time.Unix(val, 0).Format(str)
}

func UnixTimeFormatDateTime(val int64) string {
	return time.Unix(val, 0).Format(Y_M_D_H_I_S)
}
