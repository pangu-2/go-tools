package tryCatch

import (
	"errors"
	"fmt"
	"testing"
)

type err1 struct {
	error
}
type err2 struct {
	error
}

func Test_try(t *testing.T) {
	//Try1
	Try(func() {
		fmt.Println("Try1 start")
		panic(err1{error: errors.New("error1")})
	}).Catch(err1{}, func(err error) {
		fmt.Println("Try1 err1 Catch:", err.Error())
	}).Catch(err2{}, func(err error) {
		fmt.Println("Try1 err2 Catch:", err.Error())
	}).Finally(func() {
		fmt.Println("Try1 done")
	})

	//Try2
	Try(func() {
		fmt.Println("Try2 start")
		panic(err2{error: errors.New("error2")})
	}).Catch(err1{}, func(err error) {
		fmt.Println("Try2 err1 Catch:", err.Error())
	}).CatchAll(func(err error) {
		fmt.Println("Try2 Catch All:", err.Error())
	}).Finally(func() {
		fmt.Println("Try2 done")
	})

	//Try3
	Try(func() {
		fmt.Println("Try3 start")
	}).Catch(err1{}, func(err error) {
		fmt.Println("Try3 err1 Catch:", err.Error())
	}).Finally(func() {
		fmt.Println("Try3 done")
	})
}