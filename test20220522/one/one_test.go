package one

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test1(t *testing.T) {

	pc, _, _, _ := runtime.Caller(0)
	name := runtime.FuncForPC(pc).Name()

	fmt.Printf("%s start: %v\n", name, time.Now())

	time.Sleep(time.Second)

	fmt.Printf("%s end  : %v\n", name, time.Now())
}

func Test2(t *testing.T) {

	pc, _, _, _ := runtime.Caller(0)
	name := runtime.FuncForPC(pc).Name()

	fmt.Printf("%s start: %v\n", name, time.Now())

	time.Sleep(time.Second)

	fmt.Printf("%s end  : %v\n", name, time.Now())
}
