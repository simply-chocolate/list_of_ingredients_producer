package utils

import "errors"

// Takes a list of ingredients and returns a list of ingredients with the same ingredients concatenated
func HandleConcatListOfIngredients(listOfIngredients []RawMaterial, salesItem SalesItem, totalQuantityForItem float64) (SalesItem, error) {

	// Do the english list of ingredients
	listOfIngredientsEN := ""
	hasError := false
	for i, ingredient := range listOfIngredients {

		// TODO: For some ingredients, we need to add the % quantity of the ingredient
		// to figure out which ingredients we need to do this with, we need to add a field in SAP, where we can write a string of raw material item codes
		needsPercent := true

		if needsPercent {
			percentOfIngredientOnProduct := FindCorrectPrecision(ingredient.Quantity / totalQuantityForItem * 100)
			if percentOfIngredientOnProduct == "0" {
				// TODO: We need to send an error through teams, if the percent is 0
				// Error : Cannot have 0% for a claimed ingredient
				hasError = true
			} else {
				ingredient.IngredientsEnglish = IncorporatePercentAmountInIngredient(ingredient.IngredientsEnglish, percentOfIngredientOnProduct)
			}
		}

		if ingredient.IngredientsEnglish == "" {
			// send an error through teams
			hasError = true
		} else {
			if i == len(listOfIngredients)-1 {
				listOfIngredientsEN += ingredient.IngredientsEnglish
			} else {
				listOfIngredientsEN += ingredient.IngredientsEnglish + ", "
			}
		}
		// TODO: If the string for the raw material ingredientslist is empty, we need to send an error through teams.
		// We should not stop the program, but we should send an error through teams, so we can fix the problem for all raw materials at once.

		// Afterwards we need to calculate which items there can be traces of but this might be another function
	}

	//fmt.Println("listOfIngredientsEN: ", listOfIngredientsEN)
	if hasError {
		return SalesItem{}, errors.New("error: there was an error in the listOfIngredientsEN for product: " + salesItem.ItemCode)
	}

	return salesItem, nil
}
