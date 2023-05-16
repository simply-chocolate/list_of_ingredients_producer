package utils

import (
	"errors"
	"list_of_ingredients_producer/sap_api_wrapper"
	"list_of_ingredients_producer/teams_notifier"
)

// Takes a list of ingredients and returns a list of ingredients with the same ingredients concatenated
func HandleConcatAllListOfIngredients(listOfIngredients []RawMaterial, salesItem sap_api_wrapper.SapApiItemsData, totalQuantityForItem float64) (sap_api_wrapper.SapApiItemsData, error) {

	RawMaterials, err := GetRawMaterialsItemDataFromSap()
	if err != nil {
		return sap_api_wrapper.SapApiItemsData{}, err
	}

	listOfIngredientsScandinavian, err := HandleConcatListOfIngredients(listOfIngredients, totalQuantityForItem, RawMaterials, "DA_SE_NO", salesItem)
	if err != nil {
		return sap_api_wrapper.SapApiItemsData{}, err
	}
	salesItem.IngredientsScandinavian = listOfIngredientsScandinavian

	listOfIngredientsEnglish, err := HandleConcatListOfIngredients(listOfIngredients, totalQuantityForItem, RawMaterials, "EN", salesItem)
	if err != nil {
		return sap_api_wrapper.SapApiItemsData{}, err
	}
	salesItem.IngredientsEnglish = listOfIngredientsEnglish

	IngredientsFinnish, err := HandleConcatListOfIngredients(listOfIngredients, totalQuantityForItem, RawMaterials, "FI", salesItem)
	if err != nil {
		return sap_api_wrapper.SapApiItemsData{}, err
	}
	salesItem.IngredientsFinnish = IngredientsFinnish

	IngredientsGerman, err := HandleConcatListOfIngredients(listOfIngredients, totalQuantityForItem, RawMaterials, "DE", salesItem)
	if err != nil {
		return sap_api_wrapper.SapApiItemsData{}, err
	}
	salesItem.IngredientsGerman = IngredientsGerman

	IngredientsDutch, err := HandleConcatListOfIngredients(listOfIngredients, totalQuantityForItem, RawMaterials, "NL", salesItem)
	if err != nil {
		return sap_api_wrapper.SapApiItemsData{}, err
	}
	salesItem.IngredientsDutch = IngredientsDutch

	IngredientsFrench, err := HandleConcatListOfIngredients(listOfIngredients, totalQuantityForItem, RawMaterials, "FR", salesItem)
	if err != nil {
		return sap_api_wrapper.SapApiItemsData{}, err
	}
	salesItem.IngredientsFrench = IngredientsFrench

	IngredientsPortuguese, err := HandleConcatListOfIngredients(listOfIngredients, totalQuantityForItem, RawMaterials, "PT", salesItem)
	if err != nil {
		return sap_api_wrapper.SapApiItemsData{}, err
	}
	salesItem.IngredientsPortuguese = IngredientsPortuguese

	IngredientsItalian, err := HandleConcatListOfIngredients(listOfIngredients, totalQuantityForItem, RawMaterials, "IT", salesItem)
	if err != nil {
		return sap_api_wrapper.SapApiItemsData{}, err
	}
	salesItem.IngredientsItalian = IngredientsItalian

	IngredientsSpanish, err := HandleConcatListOfIngredients(listOfIngredients, totalQuantityForItem, RawMaterials, "ES", salesItem)
	if err != nil {
		return sap_api_wrapper.SapApiItemsData{}, err
	}
	salesItem.IngredientsSpanish = IngredientsSpanish

	return salesItem, nil
}

// TODO: Create a meaningful description for this function
// Takes a list of ingredients and returns a list of ingredients with the same ingredients concatenated
func HandleConcatListOfIngredients(ingredientsOnProduct []RawMaterial, totalQuantity float64, allRawMaterialsMap map[string]map[string]string, languageCode string, salesItem sap_api_wrapper.SapApiItemsData) (string, error) {

	listOfIngredients := ""
	hasError := false
	containmentMap := make(map[string]string)

	for i, ingredient := range ingredientsOnProduct {

		ingredientFromMap, exists := allRawMaterialsMap[ingredient.ItemCode]
		if !exists {
			teams_notifier.SendRawMaterialNotFoundErrorToTeams(salesItem.ItemCode, ingredient.ItemCode)
			hasError = true
		}

		containmentMap = FindAllAllergenContaminations(containmentMap, ingredientFromMap)

		if ingredientFromMap[languageCode] == "" {
			teams_notifier.SendMissingIngredientsToTeams(salesItem.ItemCode, ingredient.ItemCode, languageCode)

		} else {

			needsPercent := CheckIfIngredientIsClaimed(salesItem.ClaimedIngredients, ingredient.ItemCode)

			if needsPercent {
				// Ignore error as its just been parsed from float to string, so it should not be able to fail
				percentOfIngredientOnProduct := FindCorrectPrecision(ingredient.Quantity / totalQuantity * 100)

				if percentOfIngredientOnProduct == "0" {
					teams_notifier.SendZeroPercentClaimErrorToTeams(salesItem.ItemCode, ingredient.ItemCode)
					hasError = true

				} else {
					ingredientFromMap[languageCode] = IncorporatePercentAmountInIngredient(ingredientFromMap[languageCode], percentOfIngredientOnProduct)
				}
			}

			if i == len(ingredientsOnProduct)-1 {
				listOfIngredients += ingredientFromMap[languageCode]
			} else {
				listOfIngredients += ingredientFromMap[languageCode] + ", "
			}
		}
	}

	if hasError {
		return "", errors.New("error: there was an error in the listOfIngredients")
	}

	listOfIngredients += ". "
	cocoaDryMatterString := getCocoaDryMatterString(ingredientsOnProduct, allRawMaterialsMap, languageCode)
	listOfIngredients += cocoaDryMatterString

	containmentMap = FindAllAllergenContaminationsSalesItem(containmentMap, salesItem)
	containmentString := createStringOfTraceContamination(containmentMap, languageCode)
	listOfIngredients += containmentString

	return listOfIngredients, nil
}
