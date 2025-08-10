package main

import "fmt"

func main() {
	fmt.Println("Pokemon Type Effectiveness Simulator - Go Edition")

	var allMons []BattleMon

	for _, combo := range sampleTypeCombos {
		mon := NewBattleMon(combo)
		allMons = append(allMons, mon)
		fmt.Printf("%+v\n", mon)
	}

	attacker := allMons[0]
	defender := allMons[1]

	for _, moveName := range attacker.Moveset1 {
		move, ok := MoveProperties[moveName]
		if !ok {
			fmt.Printf("  Move %s not found in MoveProperties\n", moveName)
			continue
		}
		dmg := CalcDamage(attacker, defender, move, true)
		fmt.Printf("  Move: %-15s Damage: %d\n", moveName, dmg)
	}
}
