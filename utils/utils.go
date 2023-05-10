package utils

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// takes a map of materials as parameter and returns a list of materials sorted with the most used materials first (highest Quantity)
// Also converts the Unit Of Measure to Grams instead of Kilograms
type RawMaterial struct {
	ItemCode       string `json:"ItemCode"`
	FatherQuantity float64
	Quantity       float64 `json:"Quantity"`
}

func SortMaterialsByQuantity(materials map[string]float64) []RawMaterial {
	// First we need to convert the map to a slice
	var materialsSlice []RawMaterial
	for key, material := range materials {
		var newRawMaterial RawMaterial
		newRawMaterial.ItemCode = key
		newRawMaterial.Quantity = material * 1000

		materialsSlice = append(materialsSlice, newRawMaterial)
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
func MapRawMaterials(billOfMaterials BillOfMaterials) (map[string]float64, float64) {
	rawMaterialsMap := make(map[string]float64)
	totalQuantity := 0.0

	for _, rawMaterial := range billOfMaterials {
		// We need to somehow get the quantity of the father item, that this raw material is a part of
		_, exists := rawMaterialsMap[rawMaterial.ItemCode]
		if !exists {
			rawMaterialsMap[rawMaterial.ItemCode] = rawMaterial.Quantity

		} else {
			rawMaterialsMap[rawMaterial.ItemCode] += rawMaterial.Quantity

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

	return ingredient
}
