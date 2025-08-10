package main

// Battle-ready Pokémon struct
type BattleMon struct {
	Type1         string
	Type2         string
	HP            int
	Atk           int
	Def           int
	SpAtk         int
	SpDef         int
	Speed         int
	Moveset1      [4]string
	Moveset2      [4]string
	ActiveMoveset int // 1 or 2, defaults to 1

	//No need for other stat stages yet
	DefStage   int // -6 to +6
	SpDefStage int // -6 to +6
}

// Constructor function
func NewBattleMon(tmm TypedMonMovesets) BattleMon {
	return BattleMon{
		Type1:         tmm.Type1,
		Type2:         tmm.Type2,
		HP:            404,
		Atk:           299,
		Def:           236,
		SpAtk:         299,
		SpDef:         236,
		Speed:         236,
		Moveset1:      tmm.Moves1,
		Moveset2:      tmm.Moves2,
		ActiveMoveset: 1,
		DefStage:      0,
		SpDefStage:    0,
	}
}
