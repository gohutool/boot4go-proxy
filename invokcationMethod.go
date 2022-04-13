package proxy4go

import "reflect"

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : invokcationMethod.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/13 10:23
* 修改历史 : 1. [2022/4/13 10:23] 创建文件 by LongYong
*/

type InvocationMethod struct {
	Name string
	Type reflect.Type
}

func (im InvocationMethod) Invoke(obj any, args []reflect.Value) []reflect.Value {
	v := reflect.ValueOf(obj).MethodByName(im.Name)

	if !v.IsValid() {
		panic("Can not found method " + im.Name + " in " + reflect.ValueOf(obj).Type().String())
	}

	return v.Call(args)
}
