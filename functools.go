package functools

import (
	"reflect"
	"runtime"
	"strings"
)

func GetName(i interface{}) string {
	a := strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), "/")
	return a[len(a)-1]
}
