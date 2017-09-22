package internal

import (
	"reflect"
)

func init() {
}

func handlerMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
