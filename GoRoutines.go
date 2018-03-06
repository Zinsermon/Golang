package main

import ("fmt"
        "time"
        "sync")

// waitgroup will wait for the each go routines to complete if add 1 in it
var waitgroup sync.WaitGroup

func say(str string){
  for i:=1;i<=3;i++ {
    fmt.Println(str," is doing work #",i)
    time.Sleep(time.Millisecond*100)
  }
  // notified done with task in done
  waitgroup.Done()
}

// go chanels
func setName(c chan string, name string){
  defer waitgroup.Done()
  c <- ("Name is : "+name)
}
func main(){

fmt.Println("1st Go 1 and Go 2 will do there work concurrently then Go 3 and Go 4 will start there work cause they are dependency")
  waitgroup.Add(1)
  go say("Go 01")
  waitgroup.Add(1)
  go say("Go 02")
  waitgroup.Wait()
  fmt.Println("\n\n Now It's Go 3 And Go 4 work time \n")
  waitgroup.Add(1)
  go say("Go 03")
  waitgroup.Add(1)
  go say("Go 04")
 waitgroup.Wait()

  // in go channel...... go routines will come to end  since need to set value method
  val := make(chan string, 10)
  waitgroup.Add(1)
  go setName(val,"Salman")
  waitgroup.Add(1)
  go setName(val,"Zinsermon")
  waitgroup.Wait() 
  // close for getting values
  close(val)

  // manual getting values


  //iterating
  for name := range  val{
    fmt.Println(name)
  }

   }
