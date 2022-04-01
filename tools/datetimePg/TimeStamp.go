package datetimePg

import (
	"errors"
	"time"
)

type TimeStamp int64

func (ts TimeStamp) MarshalJSON() ([]byte, error) {
	t := time.Unix(int64(ts), 0)
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	b := make([]byte, 0, len(Y_M_D_H_I_S))
	b = append(b, '"')
	b = t.AppendFormat(b, Y_M_D_H_I_S)
	b = append(b, '"')
	return b, nil
}

func (ts *TimeStamp) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	parseTime, err := time.Parse(`"`+Y_M_D_H_I_S+`"`, string(data))
	if err != nil {
		return err
	}
	*ts = TimeStamp(parseTime.Unix())
	return nil
}
