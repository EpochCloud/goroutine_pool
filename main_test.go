package main

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
	work.Do(do1)
	result :=<- t1
	if result <= 0 {
		t.Error("do is null")
	}
}
