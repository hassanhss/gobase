package concurrents

import (
	"fmt"
	"testing"
)

/*
一、goroutine
	goroutine是Go并行设计的核心。goroutine说到底其实是协程，但它比线程要小，十几个协程在底层体现可能就是5~6个线程
	Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。

	goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。

二、Channel
	goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。那么goroutine之间如何进行数据的通信呢，Go提供了一个很好的通信机制channel。
	channel可以与Unix shell 中的双向管道做类比：可以通过它发送或者接收值。这些值只能是特定的类型：channel类型。定义一个channel时，也需要
	定义发送到channel的值的类型。注意，必须使用make 创建channel：
	eg：
		ci := make(chan int)
		cs := make(chan string)
		cf := make(chan interface{})

	channel通过操作符<-来接收和发送数据
	ch <- v    // 发送v到channel ch.
	v := <-ch  // 从ch中接收数据，并赋值给v
*/
func TestSay(t *testing.T) {
	go say("world")
	say("hello")
}

/*
默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。
所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲
channel是在多个goroutine之间同步很棒的工具
*/
func TestSum(t *testing.T) {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}
