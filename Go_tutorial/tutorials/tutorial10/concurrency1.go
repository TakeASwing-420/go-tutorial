package main
//Go routines are like normal functions but run on separate threads... by default 
//you have only the main go routine.
//But if you want, you can achieve multithreading with go routines and channels.
import (
  "fmt"
  "time"
)

func attack() {/* Output sample: 

    ~/.../cmd/tutorial10$ go run concurrency.go
    Throwing stars at Deep
    Throwing stars at Mondal
    Throwing stars at deep
    Time taken: 1.050931ms
    ~/.../cmd/tutorial10$ go run concurrency.go
    # command-line-arguments
    ./concurrency.go:16:3: smokesignal declared and not used
    ~/.../cmd/tutorial10$ go run concurrency.go
    Throwing stars at Deep
    Throwing stars at deep
    Throwing stars at Mondal
    Time taken: 106.504µs
   */  
  
  start := time.Now()

  defer func() {
    fmt.Println("Time taken:", time.Since(start))
  }()

  enemyNinjas := []string{"Deep", "deep", "Mondal"}
  smokesignal := make(chan bool, len(enemyNinjas))

  for _, _enemy := range enemyNinjas {
     // Create a local variable to capture the value of enemy for each iteration
     go func(_enemy string) {
      fmt.Println("Throwing stars at", _enemy)
      smokesignal <- true
    }(_enemy)
  }

  // Wait for all goroutines to complete
  for i := 0; i < len(enemyNinjas); i++ {
    <-smokesignal
  }
}

func main() {
  // stopChan := make(chan bool)
  // doneChan := make(chan bool)

  // go worker(stopChan, doneChan)

  // // Let the worker run for a while
  // time.Sleep(2 * time.Second)

  // // Signal the goroutine to stop
  // stopChan <- true

  // // Wait for the goroutine to confirm it has stopped
  // <-doneChan

  // close(doneChan)

  // fmt.Println("Goroutine has been stopped")
  bolvai()
}
