package utils

import (
	"list_of_ingredients_producer/sap_api_wrapper"
	"strings"
)

func FindDifferences(oldItem sap_api_wrapper.SapApiItemsData, newItem sap_api_wrapper.SapApiItemsData) string {
	differencesString := ""

	if oldItem.EnergyInKcal != strings.Replace(newItem.EnergyInKcal, ".", ",", -1) {
		differencesString += "**EnergyInKcal** changed from _" + oldItem.EnergyInKcal + "_ to _" + strings.Replace(newItem.EnergyInKcal, ".", ",", -1) + "_<BR>"
	}
	if oldItem.EnergyInkJ != strings.Replace(newItem.EnergyInkJ, ".", ",", -1) {
		differencesString += "**EnergyInkJ** changed from _" + oldItem.EnergyInkJ + "_ to _" + strings.Replace(newItem.EnergyInkJ, ".", ",", -1) + "_<BR>"
	}
	if oldItem.NutritionalFatValue != strings.Replace(newItem.NutritionalFatValue, ".", ",", -1) {
		differencesString += "**Fat** changed from _" + oldItem.NutritionalFatValue + "_ to _" + strings.Replace(newItem.NutritionalFatValue, ".", ",", -1) + "_<BR>"
	}
	if oldItem.NutritionalFattyAcidsValue != strings.Replace(newItem.NutritionalFattyAcidsValue, ".", ",", -1) {
		differencesString += "**FattyAcids** changed from _" + oldItem.NutritionalFattyAcidsValue + "_ to _" + strings.Replace(newItem.NutritionalFattyAcidsValue, ".", ",", -1) + "_<BR>"
	}
	if oldItem.NutritionalCarboHydratesValue != strings.Replace(newItem.NutritionalCarboHydratesValue, ".", ",", -1) {
		differencesString += "**CarboHydrates** changed from _" + oldItem.NutritionalCarboHydratesValue + "_ to _" + strings.Replace(newItem.NutritionalCarboHydratesValue, ".", ",", -1) + "_<BR>"
	}
	if oldItem.NutritionalSugarValue != strings.Replace(newItem.NutritionalSugarValue, ".", ",", -1) {
		differencesString += "**Sugar** changed from _" + oldItem.NutritionalSugarValue + "_ to _" + strings.Replace(newItem.NutritionalSugarValue, ".", ",", -1) + "_<BR>"
	}
	if oldItem.NutritionalProteinValue != strings.Replace(newItem.NutritionalProteinValue, ".", ",", -1) {
		differencesString += "**Protein** changed from _" + oldItem.NutritionalProteinValue + "_ to _" + strings.Replace(newItem.NutritionalProteinValue, ".", ",", -1) + "_<BR>"
	}
	if oldItem.NutritionalSaltValue != strings.Replace(newItem.NutritionalSaltValue, ".", ",", -1) {
		differencesString += "**Salt** changed from _" + oldItem.NutritionalSaltValue + "_ to _" + strings.Replace(newItem.NutritionalSaltValue, ".", ",", -1) + "_<BR>"
	}
	if oldItem.IngredientsDanish != newItem.IngredientsDanish {
		differencesString += "**IngredientsDanish** _Before:_<BR>" + oldItem.IngredientsDanish + "<BR>_After:_<BR>" + newItem.IngredientsDanish + "<BR>"
	}
	if oldItem.IngredientsDutch != newItem.IngredientsDutch {
		differencesString += "**IngredientsDutch** _Before:_<BR>" + oldItem.IngredientsDutch + "<BR>_After:_<BR>" + newItem.IngredientsDutch + "<BR>"
	}
	if oldItem.IngredientsEnglish != newItem.IngredientsEnglish {
		differencesString += "**IngredientsEnglish** _Before:_<BR>" + oldItem.IngredientsEnglish + "<BR>_After:_<BR>" + newItem.IngredientsEnglish + "<BR>"
	}
	if oldItem.IngredientsFinnish != newItem.IngredientsFinnish {
		differencesString += "**IngredientsFinnish** _Before:_<BR>" + oldItem.IngredientsFinnish + "<BR>_After:_<BR>" + newItem.IngredientsFinnish + "<BR>"
	}
	if oldItem.IngredientsFrench != newItem.IngredientsFrench {
		differencesString += "**IngredientsFrench** _Before:_<BR>" + oldItem.IngredientsFrench + "<BR>_After:_<BR>" + newItem.IngredientsFrench + "<BR>"
	}
	if oldItem.IngredientsGerman != newItem.IngredientsGerman {
		differencesString += "**IngredientsGerman** _Before:_<BR>" + oldItem.IngredientsGerman + "<BR>_After:_<BR>" + newItem.IngredientsGerman + "<BR>"
	}
	if oldItem.IngredientsItalian != newItem.IngredientsItalian {
		differencesString += "**IngredientsItalian** _Before:_<BR>" + oldItem.IngredientsItalian + "<BR>_After:_<BR>" + newItem.IngredientsItalian + "<BR>"
	}
	if oldItem.IngredientsPortuguese != newItem.IngredientsPortuguese {
		differencesString += "**IngredientsPortuguese** _Before:_<BR>" + oldItem.IngredientsPortuguese + "<BR>_After:_<BR>" + newItem.IngredientsPortuguese + "<BR>"
	}
	if oldItem.IngredientsSpanish != newItem.IngredientsSpanish {
		differencesString += "**IngredientsSpanish** _Before:_<BR>" + oldItem.IngredientsSpanish + "<BR>_After:_<BR>" + newItem.IngredientsSpanish + "<BR>"
	}
	if oldItem.IngredientsScandinavian != newItem.IngredientsScandinavian {
		differencesString += "**IngredientsScandinavian** _Before:_<BR>" + oldItem.IngredientsScandinavian + "<BR>_After:_<BR>" + newItem.IngredientsScandinavian + "<BR>"
	}

	return differencesString
}
