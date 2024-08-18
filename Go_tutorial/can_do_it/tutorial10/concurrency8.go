package main

import (
  "fmt"
  "time"
  "sync"
)
var (
  lock sync.Mutex
  rwlock sync.RWMutex
  count int
)
func hello(){
  // basic()
  readWrite()
}

func readWrite(){
  go read()
  go write()
  time.Sleep(2* time.Second) 
}

func read(){
rwlock.RLock()
defer rwlock.RUnlock()
  
fmt.Println("Reading")
time.Sleep(1*time.Second)
fmt.Println("Stopping")
}

func write(){
  rwlock.RLock()
  defer rwlock.RUnlock()

  fmt.Println("Writing")
  time.Sleep(1*time.Second)
  fmt.Println("Stopping")
}

func Increment(){
  lock.Lock()
  count++;
  lock.Unlock()
}

func basic(){
  Iterations:=1000
  for i:= 0; i< Iterations ; i++{
    go Increment()
  }
  time.Sleep(2* time.Second) 
  fmt.Println("Result",count)   
}