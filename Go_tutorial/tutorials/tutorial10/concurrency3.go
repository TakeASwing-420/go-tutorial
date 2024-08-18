package main

import (
  "fmt"
)

func Fggh() {
  ch1 := make(chan string)
  ch2 := make(chan string)

  go func() {
    ch1 <- "Message from ch1"
    close(ch1)
  }()

  go func() {
    ch2 <- "Message from ch2"
    close(ch2)
  }()

  for i := 0; i < 2; i++ {  //Outer loop decides how many times the select block will randomly pick channel calls
    select {
    case msg, open := <-ch1:
      if open {
        fmt.Println(msg)
      } else {
        fmt.Println("ch1 closed")
      }
    case msg, open := <-ch2:
      if open {
        fmt.Println(msg)
      } else {
        fmt.Println("ch2 closed")
      }
    }
  }
}
