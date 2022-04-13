package examples

import (
	. "github.com/gohutool/boot4go-proxy"
	"reflect"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : beanfactory_test.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/13 11:23
* 修改历史 : 1. [2022/4/13 11:23] 创建文件 by LongYong
*/

type SimpleBeanFactory struct {
}

func (sbf SimpleBeanFactory) NewInstance(itf any, target any) any {
	proxy := InvocationProxy.NewProxyInstance(itf, func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value {
		return method.Invoke(target, args)
	})

	return proxy
}
