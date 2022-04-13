package examples

import (
	"fmt"
	. "github.com/gohutool/boot4go-proxy"
	"reflect"
	"testing"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : example_test.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/13 10:52
* 修改历史 : 1. [2022/4/13 10:52] 创建文件 by LongYong
*/

type Hello struct {
	SayHello func() string
	SetWord  func(word string)
}

type HelloWorld struct {
	Word string
}

func (h *HelloWorld) SayHello() string {
	return h.Word
}

func (h *HelloWorld) SetWord(word string) {
	h.Word = word
}

func TestAOPProxy(t *testing.T) {
	proxy := InvocationProxy.NewProxyInstance(&Hello{}, func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf("This is a proxy function")}
	}).(*Hello)

	fmt.Println(proxy.SayHello())
}

func TestAOPProxyWithClass(t *testing.T) {
	impl := &HelloWorld{Word: "Hello world"}

	proxy := InvocationProxy.NewProxyInstance(&Hello{}, func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value {
		return method.Invoke(impl, args)
	}).(*Hello)

	proxy.SetWord("This is a proxy by HelloWorld")
	fmt.Println(proxy.SayHello())
}
