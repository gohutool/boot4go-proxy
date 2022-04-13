package proxy4go

import (
	"github.com/gohutool/log4go"
	"reflect"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : proxy4go.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/13 10:21
* 修改历史 : 1. [2022/4/13 10:21] 创建文件 by LongYong
*/

type InvocationHandler func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value

var logger = log4go.LoggerManager.GetLogger("gohutool.boot4go.proxy")

func newProxyInstance(itf any, handler InvocationHandler) any {

	t := reflect.TypeOf(itf)

	if t.Kind() != reflect.Ptr {
		panic("Need a pointer of interface struct")
	}

	if t.Elem().Kind() != reflect.Struct {
		panic("Need a pointer of interface struct")
	}

	t = t.Elem()
	ot := reflect.ValueOf(itf).Elem()
	n := ot.NumField()

	for idx := 0; idx < n; idx++ {
		f := t.Field(idx)
		of := ot.Field(idx)

		if f.Type.Kind() == reflect.Func {

			if !of.CanSet() {
				logger.Debug("field %v is a readonly func, ignore proxy", f.Name)
				continue
			}

			logger.Debug("field %v is a func, proxy now", f.Name)

			target := reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
				//param := make([]any, 0, len(args))
				//
				//for _, o := range args {
				//	param = append(param, o.Interface())
				//}

				method := InvocationMethod{Name: f.Name, Type: f.Type}

				rtn := handler(itf, method, args)

				//if rtn == nil {
				//	return []reflect.Value{}
				//}
				//
				//rtnV := make([]reflect.Value, 0, len(rtn))
				//
				//for _, o := range rtn {
				//	rtnV = append(rtnV, reflect.ValueOf(o))
				//}

				return rtn
			})
			of.Set(target)
		} else {
			logger.Debug("field %v is not a func, ignore proxy", f.Name)
		}
	}

	return itf
}
