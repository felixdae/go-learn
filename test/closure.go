package main

import "fmt"

func test_closure(){
    a := func (op int) (func()){
        i := 1
        if op == 1 {
            return func(){
                i += 1
                fmt.Println(i)
            }
        } else {
            return func(){
                i += 4
                fmt.Println(i)
            }
        }
    }
    x := a(1)
    y := a(2)
    x()
    y()
    x()
}
