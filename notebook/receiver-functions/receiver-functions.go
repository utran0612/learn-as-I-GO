//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct{
	health int
	maxHealth int
	energy int
	maxEnergy int
	name string
}

func (player *Player) modifyHealth(amount int){
	player.health += amount
	if (player.health > player.maxHealth){
		fmt.Println("Total health exceeds maximum health.")
		player.health -= amount
	}else {
		fmt.Println("Modify health successfully. Current health is",player.health)
	}

}

func (player *Player) modifyEnergy(amount int){
	player.energy += amount
	if (player.energy > player.maxEnergy){
		fmt.Println("Total energy exceeds maximum health.")
		player.health -= amount
	}else {
		fmt.Println("Modify energy successfully. Current energy is",player.energy)
	}
}

func main(){
	player := Player{0,100,0,120,"Hercules"}

	player.modifyEnergy(5)

	player.modifyHealth(90)
}