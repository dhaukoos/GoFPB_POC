package main

import (
    "context"
    "fmt"
    "math"
)
import "time"

/*
* A simple FBP example with a graph representing the Pythagorean
* formula with two emitter nodes (type int->), two squared nodes (int->int)
* a dual input and single output addition node (int=>int),
* a single input and single output square root node (int->double)
* and a single input receiver node (->double).
*/

func squared(
    inChannel <-chan int,
    outChannel chan<- int,
    ctx context.Context) {
    goProcessor(inChannel, outChannel, func(a int) int { return a * a }, ctx)
}

func squareRoot(
    inChannel <-chan int,
    outChannel chan<- float64,
    ctx context.Context) {
    goProcessor(inChannel, outChannel, func(a int) float64 { return math.Sqrt(float64(a)) }, ctx)
}

func addition(
    inChannel1 <-chan int,
    inChannel2 <-chan int,
    outChannel chan<- int,
    ctx context.Context) {
    go2inProcessor(inChannel1, inChannel2, outChannel, func(a int, b int) int { return a + b }, ctx)
}

func HelloGoFBP3() {
    chInA := make(chan int, 2)
    chInB := make(chan int, 2)
    chAtoPlus := make(chan int, 2)
    chBtoPlus := make(chan int, 2)
	chPlusToSqrt := make(chan int, 2)
    channelOutC := make(chan float64, 2)
    ctx, cancel := context.WithCancel(context.Background())

    goEmitter(chInA, ctx)
	goEmitter(chInB, ctx)
    squared(chInA, chAtoPlus, ctx)
    squared(chInB, chBtoPlus, ctx)
    addition(chAtoPlus, chBtoPlus, chPlusToSqrt, ctx)
    squareRoot(chPlusToSqrt, channelOutC, ctx)
    goReceiver(channelOutC, func(d float64) { fmt.Printf("Received here %v\n", d) }, ctx)

    time.Sleep(200 * time.Millisecond)

    // Cancel all goroutines
    fmt.Printf("Canceling received Jobs\n")
    cancel()

    // Wait for the worker to finish (optional)
    time.Sleep(time.Second)
}