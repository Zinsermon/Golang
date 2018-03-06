package main

import "fmt"

func foo() {
  defer fmt.Println("Done!")
  defer fmt.Println("Are we done yet?")
  fmt.Println("Doing some stuff, who knows what?")
}

// act as 1st-last-out
func bar() {
  for i := 0; i < 5; i++ {
  	defer fmt.Println(i)
  }
}
func main() {
  fmt.Println("defer act as 1st check function if working  then do own work , Act as FILO\n")
  foo()
  fmt.Println("\n\ndefer act as 1st check function if working  then do own work , Act as FILO\n")
  bar()

}
