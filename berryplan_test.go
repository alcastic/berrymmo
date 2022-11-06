package main

import "testing"

func TestDefinePlan(t *testing.T) {
	var seedOne, seedTwo Seed = "one", "two"
	berry := &Berry{
		name: "berry",
		recipes: []Recipe{
			{
				seedOne: 1,
				seedTwo: 1,
			},
			{
				seedOne: 0,
				seedTwo: 3,
			},
		},
	}
	var inventory = map[Seed]int{
		seedOne: 1,
		seedTwo: 5,
	}

	plan := DefinePlan(
		berry,
		inventory,
	)

	if plan.steps[0].total != 1 {
		t.Error("first step must be perform once")
	}
	if plan.steps[1].total != 1 {
		t.Error("second step must be perform once")
	}

	if plan.leftoverInventory[seedOne] != 0 {
		t.Error("leftovers seedOne must be 0")
	}
	if plan.leftoverInventory[seedTwo] != 1 {
		t.Error("leftovers seedTwo must be 1")
	}
}
