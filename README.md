# goroutine pool

##  use

```
//Initialize the number of pools
work:= Pool(10)
//Create a pool
work.NewPool()  
//calling
work.Do(Fn(do))  
//Back into the pool
work.Put(r)
```



##  use the sample programs

```
package main

import (
	"time"
	"fmt"
	"runtime"
)


//this is an instance of goroutine pool to use
func do(args ...interface{})interface{} {
	time.Sleep(2*time.Second)
	fmt.Println("work doing")

	//Here is the return value of the get function
	args[0].([]interface{})[0].(chan interface{}) <- 1
	return 1
}

func example(){
	fmt.Println("example")
}

func main(){
	r := make(chan interface{},3)  //Create a return channel
	work:= Pool(10)
	work.NewPool(r)
	fmt.Println("first",runtime.NumGoroutine())
	work.Do(Fn(do))
	work.Put(r)
	work.Do(Fn(do))
	example()
	fmt.Println(<-r)
	close(r)
	work.Close()         //close goroutine pool

	time.Sleep(5*time.Second)
}

```

