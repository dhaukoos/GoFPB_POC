package main

import (
    "context"
    "fmt"
)
import "time"

/*
* A minimalist FBP example with a single output emitter node (type int->),
* a dual input and single output processor node (int=>int),
* and a single input receiver node (->int).
*/

func HelloGoFBP2() {
    channel1 := make(chan int, 2)
    channel2 := make(chan int, 2)
	channel3 := make(chan int, 2)
    ctx, cancel := context.WithCancel(context.Background())

    goEmitter(channel1, ctx)
	goEmitter(channel2, ctx)
    go2inProcessor(channel1, channel2, channel3, func(a int, b int) int { return a + b }, ctx)
    goReceiver(channel3, func(i int) { fmt.Printf("Received here %v\n", i) }, ctx)

    time.Sleep(200 * time.Millisecond)

    // Cancel all goroutines
    fmt.Printf("Canceling emitJob\n")
    fmt.Printf("Canceling processJob\n")
    fmt.Printf("Canceling receiveJob\n")
    cancel()

    // Wait for the worker to finish (optional)
    time.Sleep(time.Second)
}