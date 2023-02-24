package main

import (
	"fmt"
)
//we create an entity interface to hold two a method wihtout implimentation

//we make the entity structs implement the methods
type entity interface{
	entityInfo() string;
	
	transformPosition()(float64,float64);

}

//two types of entitites in a game

//both have a x and a y position, and a entity name
type player struct{
	playerName string
	xPosition float64;
	yPosition float64;

}

type mob struct{
	mobType string
	xPosition float64;
	yPosition float64;

}

//print the entyties location 
func (p player) entityInfo () string{

	res := fmt.Sprintf("Player %s is located at X posion: %f, Y position: %f ", p.playerName,p.xPosition,p.yPosition)
	
	return res;

}

func (p player) transformPosition () (float64,float64){
	return p.xPosition + 10, p.yPosition - 4 ;

}
 



func (m mob) entityInfo () string{
	res := fmt.Sprintf("\nMob %s is located at X posion: %f, Y position: %f", m.mobType,m.xPosition,m.yPosition)
	
	return res;
}


func (m mob) transformPosition () (float64,float64){
	return m.xPosition + 13, m.yPosition +9 ;

}
 


func printEntityInfo(e entity){
	res := e.entityInfo()
	fmt.Println(res)

	x1,y1 := e.transformPosition()
	fmt.Println("Entities new position :", x1,", ",y1)

}


func main(){
	
	player1 := player{
		playerName: "Aldric",
		xPosition: 3,
		yPosition: 23,
		
	}

	mob1 := mob{
		mobType:"Cow",
		xPosition: 31,
		yPosition: 2,
	}

	printEntityInfo(player1)
	printEntityInfo(mob1)

}