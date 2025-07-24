/*
In Go, a goroutine is a lightweight thread managed by the Go runtime. Here’s a structured deep‑dive—just like the generics write‑up—covering definition, syntax, examples, benefits, and advanced use:
A goroutine is any function or method invoked with the go keyword. The Go scheduler multiplexes all goroutines onto a small number of OS threads, making them extremely cheap to create and switch between.

-> e.g :   go someFunction(arg1, arg2)
That call returns immediately, launching someFunction asynchronously.

The new goroutine runs concurrently with the caller.

No return value: you cannot directly capture a result—use channels or other sync to communicate back.

*/

package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Running task", id)
}


func main() {
   for i := 0; i <=10 ; i++ {
	go task(i)

   }


   time.Sleep(time.Second*2) // to make the main function sleep for 2 seconds

}