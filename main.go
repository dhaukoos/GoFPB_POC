package main

import (
    "context"
    "fmt"
)
import "time"

func main() {
	println("Hello World")
    HelloGoFBP2()
}

/*
* A minimalist FBP example with a single output emitter node (type int->),
* a single input and output processor node (int->int),
* and a single input receiver node (->int).
*/

func HelloGoFBP() {
    channel1 := make(chan int, 2)
    channel2 := make(chan int, 2)
    ctx, cancel := context.WithCancel(context.Background())

    goEmitter(channel1, ctx)
    goProcessor(channel1, channel2, func(i int) int { return i * i }, ctx)
    goReceiver(channel2, func(i int) { fmt.Printf("Received here %v\n", i) }, ctx)

    time.Sleep(200 * time.Millisecond)

    // Cancel all goroutines
    fmt.Printf("Canceling emitJob\n")
    fmt.Printf("Canceling processJob\n")
    fmt.Printf("Canceling receiveJob\n")
    cancel()

    // Wait for the worker to finish (optional)
    time.Sleep(time.Second)
}