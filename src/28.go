package main

import (
    "fmt"
)

func main() {
    var age int = 18
    if age >= 18 {
        fmt.Println("You are old enough to attend school.")
    } else {
        fmt.Println("Sorry, you need to be at least 18 years old to attend school.")
    }
}
