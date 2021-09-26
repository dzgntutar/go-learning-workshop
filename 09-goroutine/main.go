package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Println("1. gorutine calistirildi")
		wg.Done()
	}()

	// tüm wait gruplar bitene kadar bekle
	wg.Wait()

	fmt.Println("İslem tamamlandi")

}
