package main

import (
	"fmt"
	"time"
	"sync"//include sync
)

//go waitgroups
//acts as a counter
//only when all go routines have finished execution, our program will continue
//this will prevent the program from finishing before the waitgroup

//concurency, is not Parallelism
//concurency is 

func printrange (wg *sync.WaitGroup , val int){
	for i:=0; i < val; i++{
		//print value passed then i
		fmt.Println(val," : ",i)
		//sleep 1 second after after printing value
		time.Sleep(time.Second * 1)
	}

	wg.Done()

}


func main(){
	//create a wait group
	var wg sync.WaitGroup
	wg.Add(2)

	go printrange(&wg, 4)
	go printrange(&wg, 2)

 
	//if program finished befre the go routine then the go routine will also not run

	//what we could do to preven this is sleep before the program.

	//time.Sleep(10*time.Second)
	wg.Wait()
}