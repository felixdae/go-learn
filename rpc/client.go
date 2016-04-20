package main

import "fmt"
import "net/rpc"
import "log"

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

func main() {
    client, err := rpc.DialHTTP("tcp", "127.0.0.1" + ":1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    // Synchronous call
    args := &Args{27,8}
    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

    // Asynchronous call
    quotient := new(Quotient)
    divCall := client.Go("Arith.Divide", args, quotient, nil)
    // var replyCall Quotient
     <-divCall.Done // will be equal to divCall
    fmt.Println(quotient.Quo, quotient.Rem)
    // check errors, print, etc.
}
