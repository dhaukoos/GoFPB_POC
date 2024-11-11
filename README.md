# GoFPB_POC

## Example 1

<img width="403" alt="Example1" src="https://github.com/user-attachments/assets/aa56b9ef-7825-4446-bcbc-3737b4d81fe4">

A minimalist FBP example with a single output emitter node (type int->),
a single input and output processor node (int->int),
and a single input receiver node (->int).

## Example 2

<img width="403" alt="Example2" src="https://github.com/user-attachments/assets/62da9654-f9db-493a-9006-7e8b02803f1b">

A minimalist FBP example with a single output emitter node (type int->),
a dual input and single output processor node (int=>int),
and a single input receiver node (->int).

## Example 3

<img width="641" alt="Example3" src="https://github.com/user-attachments/assets/c6145e26-2d0c-4353-8565-0c6a7c48b7c8">

A simple FBP example with a graph representing the Pythagorean
formula with two emitter nodes (type int->), two squared nodes (int->int)
a dual input and single output addition node (int=>int),
a single input and single output square root node (int->double)
and a single input receiver node (->double).

## Example 4

<img width="680" alt="Example4" src="https://github.com/user-attachments/assets/34635d6a-386f-40b1-89fa-00a1d8f3e04f">

A simple FBP example with a subGraph representing the Pythagorean
formula. The non-io nodes of the previous graph are encapsulated
in a subGraph function, with passThru nodes providing the ports
into and out of the subGraph.

## Example 5

<img width="842" alt="Example5" src="https://github.com/user-attachments/assets/bd9defe6-23cf-4bd0-92d5-3dec3adb708b">

A simple FBP example with a subGraph representing the Pythagorean
formula. The differences from the previous graph include a single emitter
node that produces a Pair data object (as opposed to a simple data type),
and a single input, dual output splitter node (Pair(int)=>int).
