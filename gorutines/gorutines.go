package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "+", n)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 3; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
