package main

import "fmt"

func test_slice(){
    a := [5]int{1,2,3,4,5}
    b := a[:]
    fmt.Println(len(a), cap(a), a)
    b = append(b, 4,4,2,1,3)
    fmt.Println(len(a), cap(a), a)
    b = append(b, b...)
    fmt.Println(len(b), cap(b), b)
}
