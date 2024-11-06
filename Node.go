package main

import (
    "context"
    "fmt"
)

// Emitter function: => output only

func emitter[T any](inputSignal T, sendChannel chan<- T) {
    sendChannel <- inputSignal
}

// Processor functions: input => process => output

func processor[T, U any](receiveChannel <-chan T, sendChannel chan<- U, process func(T) U) {
    for input := range receiveChannel {
        output := process(input)
        fmt.Printf("Processing %v ==> %v\n", input, output)
        sendChannel <- output
    }
}

func twoInputAndProcessor[T, U, V any](
    inChannel1 <-chan T,
    inChannel2 <-chan U,
    sendChannel chan<- V,
    process func(T, U) V) {
        for {
            select {
                case in1, ok := <-inChannel1:
                    if !ok {
                        fmt.Println("inChannel1 is closed")
                        return
                    }
                    in2, ok := <-inChannel2
                    if !ok {
                        fmt.Println("inChannel2 is closed")
                        return
                    }
                    output := process(in1, in2)
                    fmt.Printf("Processing %v, %v ==> %v\n", in1, in2, output)
                    sendChannel <- output
                default:
                    // Wait for input from either channel
            }
        }
}
type Pair[U, V any] struct {
    a U
    b V
}
func splitPair[U, V any](inputChannel <-chan Pair[U, V], outputChannelU chan<- U, outputChannelV chan<- V) {
    for pair := range inputChannel {
        outputChannelU <- pair.a
        outputChannelV <- pair.b
    }
}

// Receiver function: input => process only

func receiver[T any](receiveChannel <-chan T, process func(T)) {
    for input := range receiveChannel {
        process(input)
        fmt.Printf("Receiving %v\n", input)
    }
}

// Node graph functions

func oneInputOneOutputGraph[T, U any](
    inChannel <-chan T,
    outChannel chan<- U,
    process func(<-chan T) chan<- U) {
        for {
            process(inChannel)
            return
        }
}

func twoInputOneOutputGraph[T, U, V any](
    inChannel1 <-chan T,
    inChannel2 <-chan U,
    outChannel chan<- V,
    process func(<-chan T, <-chan U) chan<- V) {
        for {
            process(inChannel1, inChannel2)
            return
        }
}

// Goroutine functions

// Emitter jobs: => output only

func goEmitter(outChannel chan<- int, ctx context.Context) {
    go func() {
        for i := 1; i <= 5; i++ {
            select {
                case <-ctx.Done():
                    fmt.Println("Worker was interrupted")
                    return
                default:
                    emitter(i, outChannel)
                    fmt.Printf("Emitting %v\n", i)
            }
        }
    }()
}

func goPairEmitter(outChannel chan<- Pair[int, int], ctx context.Context) {
    go func() {
        for i := 1; i <= 5; i++ {
            select {
                case <-ctx.Done():
                    fmt.Println("Worker was interrupted")
                    return
                default:
                    sides := Pair[int, int]{i, i}
                    emitter(sides , outChannel)
                    fmt.Printf("Emitting %v\n", i)
            }
        }
    }()
}

// Processor jobs: input => process => output

func goProcessor[T, U any](
    inChannel <-chan T,
    outChannel chan<- U,
    process func(T) U,
    ctx context.Context) {
    go func() {
        for {
            select {
                case <-ctx.Done():
                    fmt.Println("Worker was interrupted")
                    return
                default:
                    processor(inChannel, outChannel, process)
            }
        }
    }()
}

func go2inProcessor[T, U, V any](
    inChannel1 <-chan T,
    inChannel2 <-chan U,
    outChannel chan<- V,
    process func(T, U) V,
    ctx context.Context) {
    go func() {
        for {
            select {
                case <-ctx.Done():
                    fmt.Println("Worker was interrupted")
                    return
                default:
                    twoInputAndProcessor(inChannel1, inChannel2, outChannel, process)
            }
        }
    }()
}

func goPairSplitter[U, V any](
    inputChannel <-chan Pair[U, V],
    outputChannelU chan<- U,
    outputChannelV chan<- V,
    ctx context.Context) {
    go func() {
        for {
            select {
                case <-ctx.Done():
                    fmt.Println("Worker was interrupted")
                    return
                default:
                    splitPair(inputChannel, outputChannelU, outputChannelV)
            }
        }
    }()
}

//func goPairSplitter0[U, V any](
//    inputChannel <-chan Pair[U, V],
//    outputChannelU chan<- U,
//    outputChannelV chan<- V) <-chan struct{} {
//    done := make(chan struct{})
//    go func() {
//        splitPair(inputChannel, outputChannelU, outputChannelV)
//        close(done)
//    }()
//    return done
//}

// PassThrough jobs  input => output

func goPassThruPort[T any](
    inChannel <-chan T,
    outChannel chan<- T,
    ctx context.Context) {
    go func() {
        for input := range inChannel {
            select {
                case <-ctx.Done():
                    fmt.Println("Worker was interrupted")
                    return
                default:
                    emitter(input, outChannel)
                    fmt.Printf("Passing through %v\n", input)
            }
        }
    }()
}

// Receiver jobs: input => process only

func goReceiver[T any](inChannel <-chan T, process func(T), ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                fmt.Println("Worker was interrupted")
                return
            default:
                receiver(inChannel, process)
            }
        }
    }()
}