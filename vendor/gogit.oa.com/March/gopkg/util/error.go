package util

import (
	"fmt"
	"runtime"
)

func MustNil(err error) {
	if err != nil {
		panic(err)
	}
}

func CatchPanic(in interface{}) error {
	if in != nil {
		const size = 64 << 10
		buf := make([]byte, size)
		buf = buf[:runtime.Stack(buf, false)]
		return fmt.Errorf("catch panic: %v\n%s", in, buf)
	}
	return nil
}
