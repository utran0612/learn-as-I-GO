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

func (player *Player) ModifyHealth(amount int) bool{
	player.health += amount
	if (player.health > player.maxHealth){
		fmt.Println("Total health exceeds maximum health.")
		player.health -= amount
		return false
	}else if(player.health < 0) {
		fmt.Println("Total health is negative")
		player.health -= amount
		return false
	}else {
		fmt.Println("Modify health successfully. Current health is",player.health)
		return true
	}

}

func (player *Player) ModifyEnergy(amount int) bool{
	player.energy += amount
	if (player.energy > player.maxEnergy){
		fmt.Println("Total energy exceeds maximum health.")
		player.health -= amount
		return false
	}else if (player.energy <0){
		fmt.Println("Total energy is negative")
		player.energy -= amount
		return false
	}else {
		fmt.Println("Modify energy successfully. Current energy is",player.energy)
		return true
	}
}

func main(){
	player := Player{0,100,0,120,"Hercules"}

	player.ModifyEnergy(5)

	player.ModifyHealth(90)
}