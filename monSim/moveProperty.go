package main

type MoveProperty struct {
	BasePower     int
	Type          string
	IsHighCrit    bool
	RecoilFactor  int
	LowersDefense bool
	Split         string // "Physical" or "Special"
	Accuracy      int    // % accuracy, 0 = never misses (e.g. Swift)
}

var MoveProperties = map[string]MoveProperty{
	"Flamethrower":    {BasePower: 90, Type: "Fire", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Special", Accuracy: 100},
	"Scorching_Sands": {BasePower: 70, Type: "Ground", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Physical", Accuracy: 100},
	"Focus_Blast":     {BasePower: 120, Type: "Fighting", IsHighCrit: false, RecoilFactor: 0, LowersDefense: true, Split: "Special", Accuracy: 70},
	"Psychic":         {BasePower: 90, Type: "Psychic", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Special", Accuracy: 100},

	"Flare_Blitz": {BasePower: 120, Type: "Fire", IsHighCrit: false, RecoilFactor: 33, LowersDefense: false, Split: "Physical", Accuracy: 100},
	"Stone_Edge":  {BasePower: 100, Type: "Rock", IsHighCrit: true, RecoilFactor: 0, LowersDefense: false, Split: "Physical", Accuracy: 80},
	"Earthquake":  {BasePower: 100, Type: "Ground", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Physical", Accuracy: 100},
	"Iron_Head":   {BasePower: 80, Type: "Steel", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Physical", Accuracy: 100},

	"Ice_Beam":    {BasePower: 90, Type: "Ice", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Special", Accuracy: 100},
	"Freeze_Dry":  {BasePower: 70, Type: "Ice", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Special", Accuracy: 100},
	"Shadow_Ball": {BasePower: 80, Type: "Ghost", IsHighCrit: false, RecoilFactor: 0, LowersDefense: true, Split: "Special", Accuracy: 100},
	"Ice_Spinner": {BasePower: 80, Type: "Ice", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Physical", Accuracy: 100},
	"Crunch":      {BasePower: 80, Type: "Dark", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Physical", Accuracy: 100},

	"Energy_Ball":   {BasePower: 90, Type: "Grass", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Special", Accuracy: 100},
	"Thunderbolt":   {BasePower: 90, Type: "Electric", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Special", Accuracy: 100},
	"Poltergeist":   {BasePower: 110, Type: "Ghost", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Physical", Accuracy: 100},
	"Thunder_Punch": {BasePower: 75, Type: "Electric", IsHighCrit: false, RecoilFactor: 0, LowersDefense: false, Split: "Physical", Accuracy: 100},
}
