package main

import (
	"fmt"
)
//we create an entity interface to hold two a method wihtout implimentation

//we make the entity structs implement the methods
type entity interface{
	entityInfo() string;

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

	res := fmt.Sprintf("Player %s is located at X posion: %f, Y position: %f \n", p.playerName,p.xPosition,p.yPosition)
	
	return res;

}

func (m mob) entityInfo () string{
	res := fmt.Sprintf("Mob %s is located at X posion: %f, Y position: %f", m.mobType,m.xPosition,m.yPosition)
	
	return res;
}


func printEntityInfo(e entity){
	res := e.entityInfo()
	fmt.Println(res)
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