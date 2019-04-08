package main

import (
	"time"
	"fmt"
	"runtime"
)


//这个是goroutine pool的实例使用
func do(args ...interface{})interface{} {
	time.Sleep(2*time.Second)
	fmt.Println("work doing")

	//这里是获取函数的返回值
	args[0].([]interface{})[0].(chan interface{}) <- 1
	return 1
}

func example(){
	fmt.Println("example")
}

func main(){
	r := make(chan interface{},3)
	work:= Pool(10)
	work.NewPool(r)
	fmt.Println("first",runtime.NumGoroutine())
	work.Do(Fn(do))
	work.Put(r)
	work.Do(Fn(do))
	example()
	fmt.Println(<-r)
	fmt.Println(<-r)
	fmt.Println(<-r)
	close(r)
	work.Close()

	time.Sleep(5*time.Second)
}
