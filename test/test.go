package main

import (
    "fmt"
)

func main (){
    v := [8]int{1,2,3,4,5,6,7,8}
    s := v[4:8]
    fmt.Println(v,s,len(s),cap(s))
    s[2]=100
    fmt.Println(s)
    var (
        a,b int
        c float64
    )
    a = 1
    b = 2
    c = 89.5
    fmt.Println(a,b,c)

    test_closure()
    test_slice()
}
