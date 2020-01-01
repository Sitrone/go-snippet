package router

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func AutoBindWrap(ctrFunc interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctrType := reflect.TypeOf(ctrFunc)
		ctrValue := reflect.ValueOf(ctrFunc)
		// 1. check
		if ctrType.Kind() != reflect.Func {
			panic("not support")
			return
		}
		numIn := ctrType.NumIn()
		if numIn != 2 {
			panic("not support")
			return
		}
		
		// 2. bind value
		ctrParams := make([]reflect.Value, numIn)
		for i := 0; i < numIn; i++ {
			pt := ctrType.In(i)
			// handle gin.Context
			if pt == reflect.TypeOf(&gin.Context{}) {
				ctrParams[i] = reflect.ValueOf(c)
				continue
			}
			// handle params
			if pt.Kind() == reflect.Ptr && pt.Elem().Kind() == reflect.Struct {
				pv := reflect.New(pt.Elem()).Interface()
				var err error
				switch c.Request.Method {
				case http.MethodGet, http.MethodDelete:
					err = c.ShouldBindQuery(pv)
				case http.MethodPost, http.MethodPut:
					err = c.ShouldBindJSON(pv)
				}
				if err != nil {
					panic(err)
					return
				}

				ctrParams[i] = reflect.ValueOf(pv)
			}
		}
		// 3. call controller
		ctrValue.Call(ctrParams)
	}
}