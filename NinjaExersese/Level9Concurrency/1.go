package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go say(&wg, "hi")
	go say(&wg, "go away")
	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("About to exit")
}

func say(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Printf("I say: %s\n", message)
}
