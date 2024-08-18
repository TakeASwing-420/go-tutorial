package main

import (
  "fmt"
  "time"
)

func worker(stopChan chan bool, doneChan chan bool) {
  for {
    select {
    case <-stopChan:
      fmt.Println("Stopping goroutine")
      doneChan <- true
      close(stopChan)
      return
    default:
      // Simulate work
      fmt.Println("Working...")
      time.Sleep(500 * time.Millisecond)
    }
  }
}

