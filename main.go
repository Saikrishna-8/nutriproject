package main

import (
	"fmt"
)

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(500),
		Sugars:              SugarGram(200),
		SaturatedFattyAcids: SaturatedFattyAcidsGram(100),
		Sodium:              SodiumMilligram(100),
		Fruits:              FruitsPercent(150),
		Fibre:               FibreGram(101),
		Protein:             ProteinGram(80),
	}, Food)
	fmt.Printf("Nutritional score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
	// Output:
	// Nutritional score: 2
	// NutriScore: B
}
