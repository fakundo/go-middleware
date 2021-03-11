package middleware

import (
	"reflect"
)

// HandlerT - type of the handler.
// In future the type will be dynamic, so we will be able to make universal middlewares for any handlers with defined type.
// https://github.com/golang/go/issues/43651
type HandlerT interface{}

// Create middleware for the handler
func Create(middleware interface{}) func(handler interface{}) HandlerT {
	mw := reflect.ValueOf(middleware)
	mwT := mw.Type()
	mwNumIn := mwT.NumIn()
	nextT := mwT.In(mwNumIn - 1) // next is always the last argument

	return func(handler interface{}) HandlerT {
		h := reflect.ValueOf(handler)
		hT := h.Type()

		wrapper := reflect.MakeFunc(hT, func(in []reflect.Value) []reflect.Value {
			next := reflect.MakeFunc(nextT, func([]reflect.Value) []reflect.Value {
				res := reflect.ValueOf(handler).Call(in)
				return res[:nextT.NumOut()]
			})

			mwParams := make([]reflect.Value, mwNumIn-1)
			for i := range mwParams {
				if i < len(in) {
					mwParams[i] = in[i]
				}
			}
			mwParams = append(mwParams, next)

			return mw.Call(mwParams)
		})

		return wrapper.Interface().(HandlerT)
	}
}

// Use multiple middlewares
func Use(f ...interface{}) HandlerT {
	if len(f) < 2 {
		return f[0]
	}
	f0 := f[0].(func(interface{}) HandlerT)
	return f0(Use(f[1:]...))
}
