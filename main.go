package main

import "fmt"

func Sum(x int, y int) int {
    return x + y
}

func main() {
    ret:=Sum(5, 5)

    fmt.Printf("hello world %d \n",ret)
}

