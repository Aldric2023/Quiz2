package main 

import (
	"fmt"
)

//channel allow us to send a recieve data from a go routine


func incrementVal(val chan int, val2 int){
	val <- val2 + 1

}


func main(){
	//create a channel
	mychan := make(chan int)

	go incrementVal(mychan, 3 )
	go incrementVal(mychan, 12)

	val1 := <- mychan
	val2 := <- mychan
	//increment 3 by 1 = 4
	//inclrment 12 by 1 = 13 
	fmt.Println("values from value 1, and value 2 : ", val1,", ", val2)
}