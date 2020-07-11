package main

import (
    "fmt"
    "github.com/baconYao/go-library/protobuf/src/simple"
)

func main() {
    doSimple()
}

func doSimple() {
    sm := simplepb.SimpleMessage{
        Id: 12345,
        IsSimple: true,
        Name: "Bacon yao is my name",
        SampleList: []int32{ 1, 3, 5, 7 },
    }
    fmt.Println(sm)

    sm.Name = "I renamed you"
    fmt.Println(sm)
    fmt.Println("The ID is: ", sm.GetId())
}
