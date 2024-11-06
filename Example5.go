package main

import (
    "context"
    "fmt"
)
import "time"

/*
* A simple FBP example with a subGraph representing the Pythagorean
* formula. The differences from the previous graph include a single emitter
* node that produces a Pair data object (as opposed to a simple data type),
* and a single input, dual output splitter node (Pair(int)=>int).
*/

func pythagoreanTheoremGraph(
        inChannel <-chan Pair[int, int],
        outChannel chan<- float64,
        ctx context.Context) {
            go func() {
                for {
                    select {
                    case <-ctx.Done():
                        fmt.Println("Worker was interrupted")
                        return
                    default:
                        oneInputOneOutputGraph(inChannel, outChannel, func(<-chan Pair[int, int]) chan<- float64 {
                            chInAB := make(chan Pair[int, int], 2)
                            chA := make(chan int, 2)
                            chB := make(chan int, 2)
                            chAtoPlus := make(chan int, 2)
                            chBtoPlus := make(chan int, 2)
                            chPlusToSqrt := make(chan int, 2)
                            channelOutC := make(chan float64, 2)

                            goPassThruPort(inChannel, chInAB, ctx)
                            goPairSplitter(chInAB, chA, chB, ctx)
                            squared(chA, chAtoPlus, ctx)
                            squared(chB, chBtoPlus, ctx)
                            addition(chAtoPlus, chBtoPlus, chPlusToSqrt, ctx)
                            squareRoot(chPlusToSqrt, channelOutC, ctx)
                            goPassThruPort(channelOutC, outChannel, ctx)
                            return channelOutC
                        })
                    }
                }
            }()
    }

func HelloGoFBP5() {
    chGraphInAB := make(chan Pair[int, int], 2)

    chGraphOutC := make(chan float64, 2)
    ctx, cancel := context.WithCancel(context.Background())

    // input ports
    goPairEmitter(chGraphInAB, ctx)

    pythagoreanTheoremGraph(
            chGraphInAB,
            chGraphOutC,
            ctx)

    // output port
    goReceiver(chGraphOutC, func(d float64) { fmt.Printf("Received here %v\n", d) }, ctx)

    time.Sleep(3500 * time.Millisecond)

    // Cancel all goroutines
    fmt.Printf("Canceling received Jobs\n")
    cancel()

    // Wait for the worker to finish (optional)
    time.Sleep(time.Second)

}
