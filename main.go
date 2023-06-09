package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"time"
)

func main() {
	//fmt.Println("Let's build our first GoLang CLI")
	rand.Seed(time.Now().UTC().UnixNano())
	dice := flag.String("d", "d6", "The type of the dice to roll. Format: dx where X is an interger Defaulr: d6")
	numRoll := flag.Int("n", 1, "The number of dice to roll. Default : 1")
	sum := flag.Bool("s", false, "Get the sum of all the dice rolls")
	advantage := flag.Bool("adv", false, "Roll the dice with advantage. Default : false")
	disadvantage := flag.Bool("dis", false, "Roll the dice with disadvantage")
	flag.Parse()

	matched, _ := regexp.Match("d\\d+", []byte(*dice)) //double\\ to escpae. //takes bytes so we have to cast string dice

	if matched {
		//diceSides := (*dice)[1:] // we have to derefernce first, *dice won't work

		rolls := rollDice(dice, numRoll)
		// we're passing the pointer, we don't have to derefernce.
		//don't deference it until you need it.. memeory saving technici
		printDice(rolls)

		if *sum {
			diceSume := sumDice(rolls)
			fmt.Printf("The sume of dice was %d\n", diceSume)
		}
		if *advantage {
			roll := rollWithadvantage(rolls)
			fmt.Printf("The roll with advantage was %d\n", roll)
		}
		if *disadvantage {
			roll := rollWithDisadvantage(rolls)
			fmt.Printf("The roll with disadvantage was %d\n", roll)
		}
	} else {
		log.Fatal("Imporaor format dice dx")
	}
}