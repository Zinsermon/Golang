package main

import ("fmt"
        "math"
        "math/rand"

)

func Foo(){
  fmt.Println("Foo Is Here")
}
func Add (x,y float64) (float64){
  return x+y
}

func main(){

 /// typing Info type 1
  var a,b float64 = 4.4 ,5.5
/// typing Info type 2
   x,y := 1.1,2.2
  result := Add(a,b)
  var result_2 float64
  result_2 = Add(x,y)
  fmt.Println(result)
  fmt.Println(result_2)
  fmt.Println("Random Numnber From 1 to 100 Is",rand.Intn(100))
  fmt.Println("Square Root of 4 is ",math.Sqrt(4))
}
