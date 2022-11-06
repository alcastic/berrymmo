package main

type Seed string
type Recipe map[Seed]int
type Inventory map[Seed]int

type Berry struct {
	name    string
	recipes []Recipe
}

type RecipeStep struct {
	total  int
	recipe Recipe
}
type Plan struct {
	steps             []RecipeStep
	leftoverInventory Inventory
}

func DefinePlan(
	berry *Berry,
	seedInInventory map[Seed]int,

) *Plan {
	result := make([]RecipeStep, len(berry.recipes))
	inventory := copyInventory(seedInInventory)

	for i, recipe := range berry.recipes {
		ok := false
		maxBerries := 10000
		for k, v := range recipe {
			if v > 0 {
				nro := inventory[k] / v
				if nro < maxBerries {
					ok = true
					maxBerries = nro
				}
			}
		}
		if !ok {
			maxBerries = 0
		}
		for k, availableSeeds := range inventory {
			inventory[k] = availableSeeds - (recipe[k] * maxBerries)
		}
		result[i] = RecipeStep{total: maxBerries, recipe: recipe}
	}

	return &Plan{
		steps:             result,
		leftoverInventory: inventory,
	}
}

func copyInventory(inventory Inventory) Inventory {
	copy := make(map[Seed]int)
	for k, v := range inventory {
		copy[k] = v
	}
	return copy
}
