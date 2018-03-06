package main

import "fmt"

func main(){

// maps are like dictionary key-value pair
// make will instantiate the map
grades := make(map[string]float64)
grades["Sam"] =50
grades["Timmy"] = 64
grades["Jess"]  =92

// display
fmt.Println(grades)

// Timmy score
TimmyGrades := grades["Timmy"]
fmt.Println("Timmy grades: ",TimmyGrades)

// delete Timming from the map
delete(grades,"Timmy")
fmt.Println(grades)

// iterating through map
for key,value := range grades{
  fmt.Println(key,":",value)
}
}
