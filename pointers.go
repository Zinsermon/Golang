package main

import ("fmt")

func main(){
  var x int = 3
  a_ := &x
  b_ := x
  fmt.Println(x)
  fmt.Println(a_)
  fmt.Println(*a_)
  fmt.Println(b_)

  *a_ = 5
  b_ = 4
  fmt.Println(a_)
  fmt.Println(*a_)
  fmt.Println(b_)
  fmt.Println(x)
  *a_ = *a_**a_
  fmt.Println(x," ",*a_)
}
