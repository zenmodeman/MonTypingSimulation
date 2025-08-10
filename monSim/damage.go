package main

import (
	"fmt"
	"log"
	"math/rand"
)

func stageMultiplier(stage int) float64 {
	// clamp stage between -6 and 6
	if stage > 6 {
		stage = 6
	}
	if stage < -6 {
		stage = -6
	}

	if stage >= 0 {
		return float64(stage+2) / 2.0
	}
	return 2.0 / float64(-stage+2)
}

func TotalTypeEffectiveness(attackingType string, defendingMon BattleMon) (float64, error) {
	innerMap, ok := TypeEffectiveness[attackingType]
	if !ok {
		return 0, fmt.Errorf("unknown attacking type: %s", attackingType)
	}

	eff1, ok1 := innerMap[defendingMon.Type1]
	if !ok1 {
		eff1 = 1.0 //typeMatchups does not define neutral matchups and assumes that if an entry isn't found, the matchup is neutral
	}

	eff2 := 1.0
	if defendingMon.Type2 != "" {
		if val, ok2 := innerMap[defendingMon.Type2]; ok2 {
			eff2 = val
		}
	}

	return eff1 * eff2, nil
}

// ComputeRNG determines whether crits and damage rolls should be mocked or randomly generated
func CalcDamage(attackingMon BattleMon, defendingMon BattleMon, attackingMove MoveProperty, computeRNG bool) int {
	levelFactor := 42 //At level 100, the level formula evaluates to this value
	power := attackingMove.BasePower
	split := attackingMove.Split
	offensive := attackingMon.Atk
	defensive := defendingMon.Def
	defensiveStage := defendingMon.DefStage

	if split == "Special" {
		offensive = attackingMon.SpAtk
		defensive = attackingMon.SpDef
		defensiveStage = defendingMon.SpDefStage
	}

	baseDamage := int(float64(levelFactor*power*offensive) / (float64(defensive) * (stageMultiplier(defensiveStage)) * 50))
	critMod := 1.0
	randomRoll := 0.93 //Roughly the average damage roll, for predictions

	//computeRNG is false when this function is called for predictions instead of actual damage.
	if computeRNG {
		//RNG for crits
		critChance := 1.0 / 24.0
		if attackingMove.IsHighCrit {
			critChance *= 3
		}
		if rand.Float64() <= critChance {
			critMod = 1.5
		}

		//RNG for damage rolls
		rollSteps := 16
		rollIndex := rand.Intn(rollSteps)
		randomRoll = 0.85 + float64(rollIndex)*0.01
	}
	damageAfterCrit := int(float64(baseDamage+2) * critMod)
	damageAfterRandomRoll := int(float64(damageAfterCrit) * randomRoll)

	stabDmg := 1.0
	if attackingMon.Type1 == attackingMove.Type || attackingMon.Type2 == attackingMove.Type {
		stabDmg = 1.5
	}
	damageAfterStab := int(float64(damageAfterRandomRoll) * stabDmg)
	effectiveness, err := TotalTypeEffectiveness(attackingMove.Type, defendingMon)
	if err != nil {
		log.Fatalf("Failed to compute effectiveness: %v", err)
	}

	damageAfterEffectiveness := int(float64(damageAfterStab) * effectiveness)

	finalDamage := damageAfterEffectiveness //This could be subject to change if I add in more damage multipliers

	return finalDamage
}
