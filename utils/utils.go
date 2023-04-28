package utils

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// takes a map of materials as parameter and returns a list of materials sorted with the most used materials first (highest Quantity)
// Also converts the Unit Of Measure to Grams instead of Kilograms
func SortMaterialsByQuantity(materials map[string]RawMaterial) []RawMaterial {
	// First we need to convert the map to a slice
	var materialsSlice []RawMaterial
	for _, material := range materials {
		material.Quantity *= 1000
		materialsSlice = append(materialsSlice, material)
	}

	sort.Slice(materialsSlice, func(i, j int) bool {
		return materialsSlice[i].Quantity > materialsSlice[j].Quantity
	})

	return materialsSlice
}

// Rounds a float to the specified precision
func RoundWithPrecision(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Round(num*output)) / output
}

// Takes a float and returns a string with the correct precision depending on the size
func FindCorrectPrecision(ingredientPercent float64) string {
	if ingredientPercent > 10.0 {
		return fmt.Sprint(math.Round(ingredientPercent))
	} else if ingredientPercent > 1.0 {
		return fmt.Sprint(RoundWithPrecision(ingredientPercent, 1))
	} else {
		return fmt.Sprint(RoundWithPrecision(ingredientPercent, 2))
	}
}

// takes a bill of materials
// return a map of raw materials and a total quantity
func MapRawMaterials(billOfMaterials BillOfMaterials) (map[string]RawMaterial, float64) {
	rawMaterialsMap := make(map[string]RawMaterial)
	totalQuantity := 0.0

	for _, rawMaterial := range billOfMaterials {
		// We need to somehow get the quantity of the father item, that this raw material is a part of
		value, exists := rawMaterialsMap[rawMaterial.ItemCode]
		if !exists {
			newRawMaterial := RawMaterial{
				ItemCode:                rawMaterial.ItemCode,
				Quantity:                rawMaterial.Quantity,
				IngredientsScandinavian: rawMaterial.IngredientsScandinavian,
				IngredientsFinnish:      rawMaterial.IngredientsFinnish,
				IngredientsEnglish:      rawMaterial.IngredientsEnglish,
				IngredientsGerman:       rawMaterial.IngredientsGerman,
				IngredientsDutch:        rawMaterial.IngredientsDutch,
				IngredientsFrench:       rawMaterial.IngredientsFrench,
				IngredientsPortuguese:   rawMaterial.IngredientsPortuguese,
				IngredientsItalian:      rawMaterial.IngredientsItalian,
				IngredientsSpanish:      rawMaterial.IngredientsSpanish,
			}
			rawMaterialsMap[rawMaterial.ItemCode] = newRawMaterial
		} else {
			value.Quantity += rawMaterial.Quantity
			rawMaterialsMap[rawMaterial.ItemCode] = value
		}
		totalQuantity += rawMaterial.Quantity
	}
	// Convert totalQuantity to grams
	totalQuantity *= 1000

	return rawMaterialsMap, totalQuantity
}

// Takes an ingredient and a percent amount
// Checks if the ingredient contains any paranthesis, if it does, it adds the percent amount before the opening paranthesis, if not, it adds the percent amount after the ingredient
func IncorporatePercentAmountInIngredient(ingredient string, percentAmount string) string {
	parenthesesIndex := strings.Index(ingredient, "(")
	if parenthesesIndex == -1 {
		ingredient = ingredient + " (" + percentAmount + "%)"
	} else {
		ingredient = ingredient[:parenthesesIndex] + "(" + percentAmount + "%)" + ingredient[parenthesesIndex:]
	}
	fmt.Println("ingredient: ", ingredient)
	return ingredient
}
