package main

import (
	"fmt"
	"time"
)

//concurency, is not Parallelism
//concurency is 

func printrange (val int){
	for i:=0; i < val; i++{
		//print value passed then i
		fmt.Println(val," : ",i)
		//sleep 1 second after after printing value
		time.Sleep(time.Second * 1)
	}
}


func main(){
	go printrange(4)
	go printrange(9)

	//if program finished befre the go routine then the go routine will also not run
	
	//what we could do to preven this is sleep before the program.

	time.Sleep(10 * time.Second)

}