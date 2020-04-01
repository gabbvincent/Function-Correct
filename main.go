package main

import (
  "fmt"
  "math/rand"
  "time"
)



func correct(){
  var a [5]string
  a[0] = "Correct!"
  a[1] = "That's Right!"
  a[2] = "Good Job!"
  a[3] = "Nice!"
  a[4] = "Yes! You got it!"

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  i := r1.Intn(5)

  fmt.Println(a[i])
}



func main() {
  
  for i := 0; i < 10; i++{
    correct()
  }
}