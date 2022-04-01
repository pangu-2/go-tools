package datetimePg

import (
	"fmt"
	"testing"
	"time"
)

type datetime2 struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	CreatedAt DateTime `json:"created_at"`
	UpdatedAt DateTime `json:"updated_at"`
}

func TestDatetime2(t *testing.T) {
	dateTime := DateTime(time.Now())
	fmt.Println(dateTime)
	fmt.Println(dateTime.String())
}
