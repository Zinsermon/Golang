package main

import "fmt"

// const

const usixteenbitmax float64 = 65535
const kmh_multiple float64 =  1.6

type Car struct {
  gas_pedal uint16 //  min: 0,      max: 65535
	brake_pedal uint16   //min: 0,      max: 65535
	steering_wheel int16 //min: -32768  max: 32768
  top_speed_kmh float64
}

/// Methods only to the struct cannot use directly + no func can access struct

/// Values Recievers
func (car *Car) kmh() float64 {
  return float64 (car.gas_pedal) *(car.top_speed_kmh/usixteenbitmax)
}

func (car *Car) mph() float64{
  return float64 (car.gas_pedal) *(car.top_speed_kmh/usixteenbitmax/kmh_multiple)
}

/// Pointer Recievers
func (car *Car) new_top_speed(new_speed float64)  {
  car.top_speed_kmh =new_speed
}

func (car *Car) new_gas_pedal(gas_pedal uint16)  {
  car.gas_pedal =gas_pedal
}

func (car *Car) new_brake_pedal(brake_pedal uint16)  {
  car.brake_pedal =brake_pedal
}

func (car *Car) new_steering_wheel(steering_wheel int16)  {
  car.steering_wheel =steering_wheel
}


func main(){
  var car Car = Car{
    gas_pedal:10,
    brake_pedal:0,
    steering_wheel:100,
    top_speed_kmh:100,
  }

  fmt.Println("Gas Pedal: ",car.gas_pedal)
  fmt.Println("Brake pedal:",car.brake_pedal)
  fmt.Println("Steering Wheel: ",car.steering_wheel)
  fmt.Println("Top Speed is",car.top_speed_kmh,"\n")

  fmt.Println("kph is:",car.kmh())
  fmt.Println("mph is:",car.mph(),"\n")

  car.new_top_speed(1000)
  car.new_gas_pedal(44)
  car.new_gas_pedal(19)
  car.new_steering_wheel(90)

  fmt.Println("Gas Pedal: ",car.gas_pedal)
  fmt.Println("Brake pedal:",car.brake_pedal)
  fmt.Println("Steering Wheel: ",car.steering_wheel)
  fmt.Println("Top Speed is",car.top_speed_kmh)

}
