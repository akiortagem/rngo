package rngo

import (
	"math/rand"
	"sort"
)

func rollNDice(nDice int, sides int) []int {
	//list to contain dice rolls
	var rolls []int

	//loop through nDice
	for i := 0; i < nDice; i++ {
		//append rolls with the new roll
		rolls = append(rolls, rand.Intn(sides)+1)
	}
	return rolls
}

func selectAdvantage(rolls []int, adv int) []int {
	// selects the highest or lowest roll from a list of rolls
	// positive adv keeps highest rolls, negative adv keeps lowest rolls
	// adv = 0 returns the original list of rolls

	//early return
	if adv == 0 {
		return rolls
	}

	//first sort the rolls in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(rolls)))

	if adv > 0 {
		//return the highest rolls
		return rolls[:adv]
	}
	// turn the negative number into a positive number
	adv = adv * -1
	// select the lowest rolls
	return rolls[len(rolls)-adv:]
}

func explode(rolls []int, explodeOn int, sides int) []int {
	// explodes the dice on specified number
	// returns the original list of rolls if explodeOn is not in the list
	// will roll the exploded dice again based on sides

	//early return for explodeOn <= 0
	if explodeOn <= 0 {
		return rolls
	}

	explodedRolls := []int{}

	//copy rolls so that we can pop and push from it
	rollsCopy := make([]int, len(rolls))

	//as long as there are still rolls in the list
	//pop the first roll and roll again
	for len(rollsCopy) > 0 {
		//pop the first roll
		roll := rollsCopy[0]
		rollsCopy = rollsCopy[1:]
		//roll again if the roll is equal to explodeOn
		if roll == explodeOn {
			explodedRolls = append(explodedRolls, rollNDice(1, sides))
		}

	//combine both list
	rolls = append(rolls, explodedRolls...)

	return rolls
}