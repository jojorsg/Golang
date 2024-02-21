package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitgroup sync.WaitGroup

	waitgroup.Add(2)

	go func() {
		escrever("Ola mundo")
		waitgroup.Done()
	}()
	
	go func() {
		escrever("Programando em Go!")
		waitgroup.Done()
	}()

	waitgroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}