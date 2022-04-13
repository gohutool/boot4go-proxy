package examples

import (
	. "github.com/gohutool/boot4go-proxy"
	"reflect"
	"testing"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : benchmark_test.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/13 11:29
* 修改历史 : 1. [2022/4/13 11:29] 创建文件 by LongYong
*/

func BenchmarkWithoutProxy(b *testing.B) {
	b.SetParallelism(1)
	hello := &HelloWorld{}
	hello.SetWord("This is a proxy by HelloWorld")

	for i := 0; i < b.N; i++ {
		hello.SayHello()
	}
}

func BenchmarkWithoutFuncProxy(b *testing.B) {
	b.SetParallelism(1)
	proxy := InvocationProxy.NewProxyInstance(&Hello{}, func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf("This is a proxy by HelloWorld")}
	}).(*Hello)

	for i := 0; i < b.N; i++ {
		proxy.SayHello()
	}
}

func BenchmarkWithProxy(b *testing.B) {
	b.SetParallelism(1)
	hello := &HelloWorld{}
	proxy := InvocationProxy.NewProxyInstance(&Hello{}, func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value {
		return method.Invoke(hello, args)
	}).(*Hello)
	proxy.SetWord("This is a proxy by HelloWorld")

	for i := 0; i < b.N; i++ {
		proxy.SayHello()
	}
}
