package main

import (
  "fmt"
  "os"
)

func main() {
  firstGreeting := os.Getenv("INPUT_FIRSTGREETING")
  secondGreeting := os.Getenv("INPUT_SECONDGREETING")
  thirdGreeting := os.Getenv("INPUT_THIRDGREETING")
  
  fmt.Println("Hello " + firstGreeting)
  fmt.Println("Hello " + secondGreeting)
  
  if thirdGreeting != "" {
    fmt.Println("Hello " + thirdGreeting)
  }
}
