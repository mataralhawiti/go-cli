package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strconv"
)

func rollDice(dice *string, times *int) []int {
	var rolls []int

	diceSides := (*dice)[1:] // we have to derefernce first, *dice won't work
	d, err := strconv.Atoi(diceSides)

	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < *times; i++ {
		rolls = append(rolls, rand.Intn(d)+1)
	}
	return rolls
}

func printDice(rolls []int) {
	for i, dice := range rolls {
		fmt.Printf("Roll %d was %d\n", i+1, dice)
	}
}

func sumDice(rolls []int) int {
	sum := 0
	for _, dice := range rolls {
		sum += dice
	}
	return sum
}

func rollWithadvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[len(rolls)-1]
}

func rollWithDisadvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[0]
}