package goroutine_pool

import (
	"testing"
	"fmt"
)

func do1(args ...interface{})(interface{}){
	fmt.Println("test1")
	args[0].([]interface{})[0].(chan int) <- 1
	return 1
}

func Test_m(t *testing.T){
	t1 := make(chan int,1)
	defer close(t1)
	work:= Pool(10)
	if work == nil {
		t.Error("work is null")
	}
	work.NewPool(t1)
	work.Do(do1).Put()
	result :=<- t1
	if result <= 0 {
		t.Error("do is null")
	}
}

/*
*use
**/

//func do(args ...interface{})interface{} {
//	time.Sleep(2*time.Second)
//	fmt.Println("work doing")
//
//	//这里是获取函数的返回值
//	args[0].([]interface{})[0].(chan interface{}) <- 1
//	return 1
//}
//func main(){
//	r := make(chan interface{},3)
//	work:= Pool(10)
//	work.NewPool(r)   //use
//	fmt.Println("first",runtime.NumGoroutine())
//	work.Do(Fn(do)).Put(r) //get and put
//	fmt.Println(<-r)       //get result
//	close(r)
//	work.Close()
//	time.Sleep(5*time.Second)
//}