package main

import (
	"fmt"
	"sync"
)

var (
	pty = [...]string{"Tommy", "kolman", "jiol", "bob"}
)

func jio() {
	var wg sync.WaitGroup

	for _, en := range pty {
    wg.Add(1)
    
		go func(enemy_ninja string) {
			atk(enemy_ninja, &wg)
			return
		}(en)

	}
	wg.Wait()
	fmt.Println("Done")

}

func atk(en string, deepwait *sync.WaitGroup) {
	fmt.Println("Hey! Attacked evil Ninja:", en)
	deepwait.Done()
}
