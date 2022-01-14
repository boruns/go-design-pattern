# 单例模式

单例模式有两种：**饿汉式** 和 **懒汉式**

## 定义

    一个类只允许创建一个对象（或者叫实例），那这个类就是单例类。这种设计模式就叫单例设计模式，简称单例模式。

## 用处

    从业务概念上，有些数据在系统中只应该保存一份，就比较适合设计成单例类。

## 单例模式的唯一性

### 进程间唯一
### 如何实现线程间唯一

    通过获取线程id，但在go中无法实现，go主要使用的协程，并且协程id没有暴露出来。

### 如何实现集群环境唯一？（多进程）

    通过外部共享存储的锁进行。例如文件，`redis`等。


## 存在问题

### 会隐藏类之间的依赖关系

    单例模式降低可读性，通过构造函数、参数传递等方式声明的类之间的依赖关系，我们通过查看函数的定义，就能很清晰的识别出来。

### 对代码扩展性不友好

    如果有一天我们需要创建多个实例就很麻烦。

### 对代码可预测性不友好

    单例类是硬编码模式，不利于`mock`。
    如果是可修改的全局变量，测试的时候还需要注意不同的测试用例对它的修改可能会有问题。（可以通过将单例类作为参数传递给需要使用的方法解决可预测性的问题）

## 如何实现

### 构造函数是private访问权限
### 考虑对象创建时的线程安全问题
### 考虑是否支持延迟加载
### 考虑getInstance的性能问题（是否有加锁等）

## 实现方式

### 饿汉式

    类加载的时候`instance`实例已经创建好了。
    `isstance`的创建是线程安全的。
    不支持延时加载。
    初始化时间可能比较长。（延迟加载不是问题，但是系统启动的时候慢一些比用户使用时慢一些更好些）

### 懒汉式

    在`getInstance`时再去创建实例。
    `instance`创建时需要加锁。
    支持延迟加载。
    不支持高并发。

### 双重检测

    在懒汉式的基础上，将方法的锁改成类级别的锁。
    相对于懒汉式的锁粒度更小，不会每次都去获取锁。


## Testing 结果

### 饿汉式

```
go test -timeout 30s -run ^TestGetHungryInstance$ github.com/boruns/go-design/test/singleton

ok  	github.com/boruns/go-design/test/singleton	0.008s



go test -benchmem -run=^$ -bench ^BenchmarkGetHungryInstanceParallel$ github.com/boruns/go-design/test/singleton

goos: darwin
goarch: amd64
pkg: github.com/boruns/go-design/test/singleton
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkGetHungryInstanceParallel-8   	1000000000	         0.2643 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/boruns/go-design/test/singleton	0.302s
```

### 懒汉式

```
go test -timeout 30s -run ^TestGetLazyInstance$ github.com/boruns/go-design/test/singleton

ok  	github.com/boruns/go-design/test/singleton	0.008s


go test -benchmem -run=^$ -bench ^BenchmarkGetLazyInstance$ github.com/boruns/go-design/test/singleton

goos: darwin
goarch: amd64
pkg: github.com/boruns/go-design/test/singleton
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkGetLazyInstance-8   	1000000000	         0.9783 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/boruns/go-design/test/singleton	1.085s
```



