
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func election() {
  ch1, ch2 := make(chan bool), make(chan bool)

  // Create a new local random number generator
  rng := rand.New(rand.NewSource(time.Now().UnixNano()))

  var Harshit, Naresh uint32

  // Start goroutines to send votes
  go func() {
    sendVotes(ch1, true, rng.Intn(100))
  }()
  go func() {
    sendVotes(ch2, true, rng.Intn(100))
  }()

  // Collect votes
  for i := 0; i < 1000; i++ {
    select {
    case <-ch1:
      Harshit++
    case <-ch2:
      Naresh++
    }
  }

  fmt.Println("Naresh:", Naresh, "Harshit:", Harshit)
}

func sendVotes(ch chan bool, vote bool, numVotes int) {
  for i := 0; i < numVotes; i++ {
    ch <- vote
  }
  close(ch)
}