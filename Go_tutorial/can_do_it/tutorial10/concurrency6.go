package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (missioncomplete bool
    r = rand.New(rand.NewSource(time.Now().Unix())))

func leonardo() {
	var wg sync.WaitGroup
  // var once sync.Once
  
	wg.Add(100)
for i:= 0; i <= 100; i++{
	go func() {
		if FoundTreasure() {
			// once.Do(markMissionComplete)
      markMissionComplete()
		}
		wg.Done()
	}()
}
	wg.Wait()
}

func markMissionComplete() {
	missioncomplete = true
  check()
}

func check() {
	if missioncomplete {
		fmt.Println("Mission complete")
	} else {
		fmt.Println("Mission is not complete")
	}
}

func FoundTreasure() bool {
	a := r.Intn(10)
	return 0 == a
}
