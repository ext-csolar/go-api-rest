package main

import (
	"fmt"
	"time"
)

func main() {
	dia := time.Now().Weekday()

    fmt.Println("hello world! running on: ", dia)
}