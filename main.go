package main

import "sync"

const url = "https://funpay.ru/chips/2/"

func main() {

	wg := sync.WaitGroup{}

	p := NewParser(url)

	wg.Add(1)
	go p.Run(&wg)

	wg.Wait()
}
