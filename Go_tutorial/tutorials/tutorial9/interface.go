package main

import (
  "fmt"
  "math/rand"
)

type Player interface {
  KickBall()
}

type footballplayer1 struct {
  stamina     int
  power       int
  healthboost int
}

type footballplayer2 struct {
  stamina int
  power   int
}

func (fp1 *footballplayer1) KickBall() { // Method
  shot := fp1.stamina + fp1.power*fp1.healthboost
  fmt.Println("I am stronger...I am kicking the ball with power:", shot)
}

func (fp2 *footballplayer2) KickBall() { // Method
  shot := fp2.stamina + fp2.power
  fmt.Println("I am kicking the ball with power:", shot)
}

func main() {
  r := rand.New(rand.NewSource(42)) //Create a local seed generator
  team := make([]Player, 11)
  for i := 0; i < len(team); i++ {
    if i == 7 {
      team[i] = &footballplayer2{
        stamina: 25,
        power:   r.Intn(100),
      }
    } else {
      team[i] = &footballplayer1{
        stamina:     90,
        power:       r.Intn(100),
        healthboost: r.Intn(75),
      }
    }
  }

  for _, player := range team {
    player.KickBall()
  }
}
