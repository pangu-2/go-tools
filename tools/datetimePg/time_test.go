package datetimePg

import (
	"fmt"
	"testing"
	"time"
)

func TestDatetime(t *testing.T) {
	//time := DateTime.Time()
	timer := DateTime(time.Now())
	fmt.Println(timer)
	fmt.Println(timer.Time())
	fmt.Println(timer.Date())
	fmt.Println(timer.DateTime())
	fmt.Println(timer.MarshalJSON())
	fmt.Println(timer.Format(Y_M_D_H_I_S))
}
