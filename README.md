# boot4go-proxy

a dynamic proxy toolkit for golang like as Proxy.newInstance like as java

![license](https://img.shields.io/badge/license-Apache--2.0-green.svg)

# Introduce
This project is a lightly dynamic proxy implementation for golang, it like as the "Proxy.newInstance" in java.
In order that you can proxy a defined interface with target class dynamically. 
As we know we can proxy a class with another class by proxy design pattern, it will be static mode and need some code to implemented it.
With the proxy4go we do implement it dynamically and easier.

# Feature
With the proxy4go we do implement it dynamically and easier.

- Proxy the source with destination class
- Proxy the source with func object


# Usage
- Add proxy4go with the following import

```
import . "github.com/gohutool/gohutool/boot4go-proxy"
```

- Define interface struct

```
type Hello struct {
    SayHello func() string
}
```

- Dynamical proxy this interface with a function

```
func TestAOPProxy(t *testing.T) {
	proxy := InvocationProxy.NewProxyInstance(&Hello{}, func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf("This is a proxy function")}
	}).(*Hello)

	fmt.Println(proxy.SayHello())
}
```

```
=== RUN   TestAOPProxy
This is a proxy function
--- PASS: TestAOPProxy (0.00s)
```

- Dynamical proxy this interface with a destination class

```
type Hello struct {
	SayHello func() string
	SetWord  func(word string)
}
```

```
type HelloWorld struct {
	Word string
}

func (h *HelloWorld) SayHello() string {
	return h.Word
}

func (h *HelloWorld) SetWord(word string) {
	h.Word = word
}
```

```
func TestAOPProxyWithClass(t *testing.T) {
	impl := &HelloWorld{Word: "Hello world"}

	proxy := InvocationProxy.NewProxyInstance(&Hello{}, func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value {
		return method.Invoke(impl, args)
	}).(*Hello)

	proxy.SetWord("This is a proxy by HelloWorld")
	fmt.Println(proxy.SayHello())
}
```

```
=== RUN   TestAOPProxyWithClass
This is a proxy by HelloWorld
--- PASS: TestAOPProxyWithClass (0.00s)
```

- With proxy4go we can implement like BeanFactory in java spring/springboot in golang

```
type SimpleBeanFactory struct {
}

func (sbf SimpleBeanFactory) NewInstance(itf any, target any) any {
	proxy := InvocationProxy.NewProxyInstance(itf, func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value {
		return method.Invoke(target, args)
	})

	return proxy
}
```

# Benchmark testing

- Testing invoking directly
```
func BenchmarkWithoutProxy(b *testing.B) {
	b.SetParallelism(1)
	hello := &HelloWorld{}
	hello.SetWord("This is a proxy by HelloWorld")

	for i := 0; i < b.N; i++ {
		hello.SayHello()
	}
}
```

```
goos: windows
goarch: amd64
pkg: github.com/gohutool/boot4go-proxy/examples
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkWithoutProxy
BenchmarkWithoutProxy-8         385053259                3.189 ns/op
PASS
```

- Testing invoking by function proxy 

```
func BenchmarkWithoutFuncProxy(b *testing.B) {
	b.SetParallelism(1)
	proxy := InvocationProxy.NewProxyInstance(&Hello{}, func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf("This is a proxy by HelloWorld")}
	}).(*Hello)

	for i := 0; i < b.N; i++ {
		proxy.SayHello()
	}
}
```

```
goos: windows
goarch: amd64
pkg: github.com/gohutool/boot4go-proxy/examples
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkWithoutFuncProxy
BenchmarkWithoutFuncProxy-8      2781796               376.8 ns/op
PASS
```

- Testing invoking by destination class proxy

```
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
```

```
goos: windows
goarch: amd64
pkg: github.com/gohutool/boot4go-proxy/examples
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkWithProxy
BenchmarkWithProxy-8      643676              1735 ns/op
PASS
```

# Different with pig and monkey
