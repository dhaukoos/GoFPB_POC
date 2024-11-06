package main

import (
    "context"
    "fmt"
)
import "time"

/*
* A simple FBP example with a subGraph representing the Pythagorean
* formula. The non-io nodes of the previous graph are encapsulated
* in a subGraph function, with passThru nodes providing the ports
* into and out of the subGraph.
*/

func pythagoreanTheoremGraph1(
        inChannel1 <-chan int,
        inChannel2 <-chan int,
        outChannel chan<- float64,
        ctx context.Context) {
            go func() {
                for {
                    select {
                    case <-ctx.Done():
                        fmt.Println("Worker was interrupted")
                        return
                    default:
                        twoInputOneOutputGraph(inChannel1, inChannel2, outChannel, func(<-chan int, <-chan int) chan<- float64 {
                            chInA := make(chan int, 2)
                            chInB := make(chan int, 2)
                            chAtoPlus := make(chan int, 2)
                            chBtoPlus := make(chan int, 2)
                            chPlusToSqrt := make(chan int, 2)
                            channelOutC := make(chan float64, 2)

                            goPassThruPort(inChannel1, chInA, ctx)
                            goPassThruPort(inChannel2, chInB, ctx)
                            squared(chInA, chAtoPlus, ctx)
                            squared(chInB, chBtoPlus, ctx)
                            addition(chAtoPlus, chBtoPlus, chPlusToSqrt, ctx)
                            squareRoot(chPlusToSqrt, channelOutC, ctx)
                            goPassThruPort(channelOutC, outChannel, ctx)
                            return channelOutC
                        })
                    }
                }
            }()
    }

func HelloGoFBP4() {
    chGraphInA := make(chan int, 2)
    chGraphInB := make(chan int, 2)

    chGraphOutC := make(chan float64, 2)
    ctx, cancel := context.WithCancel(context.Background())

    // input ports
    goEmitter(chGraphInA, ctx)
    goEmitter(chGraphInB, ctx)

    pythagoreanTheoremGraph1(
            chGraphInA,
            chGraphInB,
            chGraphOutC,
            ctx)

    // output port
    goReceiver(chGraphOutC, func(d float64) { fmt.Printf("Received here %v\n", d) }, ctx)

    time.Sleep(4000 * time.Millisecond)

    // Cancel all goroutines
    fmt.Printf("Canceling received Jobs\n")
    cancel()

    // Wait for the worker to finish (optional)
    time.Sleep(time.Second)

}
