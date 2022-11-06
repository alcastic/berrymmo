package main

import (
	"fmt"
)

const sweetSeed Seed = "sweet seed"
const verySweetSeed Seed = "very sweet seed"

var MelocBerry *Berry

func init() {
	MelocBerry = &Berry{
		name: "meloc",
		recipes: []Recipe{
			{
				verySweetSeed: 1,
				sweetSeed:     1,
			},
			{
				sweetSeed: 3,
			},
		},
	}
}

func main() {
	fmt.Println("start")

	var seedInInventory = map[Seed]int{
		verySweetSeed: 2,
		sweetSeed:     1,
	}

	plan := DefinePlan(
		MelocBerry,
		seedInInventory,
	)
	fmt.Println(*plan)

	fmt.Println("finish")
}
