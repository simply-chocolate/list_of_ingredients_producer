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
		differencesString += "**IngredientsDanish** changed from <BR>_" + oldItem.IngredientsDanish + "_ to <BR>_" + newItem.IngredientsDanish + "_<BR>"
	}
	if oldItem.IngredientsDutch != newItem.IngredientsDutch {
		differencesString += "**IngredientsDutch** changed from <BR>_" + oldItem.IngredientsDutch + "_ to <BR>_" + newItem.IngredientsDutch + "_<BR>"
	}
	if oldItem.IngredientsEnglish != newItem.IngredientsEnglish {
		differencesString += "**IngredientsEnglish** changed from <BR>_" + oldItem.IngredientsEnglish + "_ to <BR>_" + newItem.IngredientsEnglish + "_<BR>"
	}
	if oldItem.IngredientsFinnish != newItem.IngredientsFinnish {
		differencesString += "**IngredientsFinnish** changed from <BR>_" + oldItem.IngredientsFinnish + "_ to <BR>_" + newItem.IngredientsFinnish + "_<BR>"
	}
	if oldItem.IngredientsFrench != newItem.IngredientsFrench {
		differencesString += "**IngredientsFrench** changed from <BR>_" + oldItem.IngredientsFrench + "_ to <BR>_" + newItem.IngredientsFrench + "_<BR>"
	}
	if oldItem.IngredientsGerman != newItem.IngredientsGerman {
		differencesString += "**IngredientsGerman** changed from <BR>_" + oldItem.IngredientsGerman + "_ to <BR>_" + newItem.IngredientsGerman + "_<BR>"
	}
	if oldItem.IngredientsItalian != newItem.IngredientsItalian {
		differencesString += "**IngredientsItalian** changed from <BR>_" + oldItem.IngredientsItalian + "_ to <BR>_" + newItem.IngredientsItalian + "_<BR>"
	}
	if oldItem.IngredientsPortuguese != newItem.IngredientsPortuguese {
		differencesString += "**IngredientsPortuguese** changed from <BR>_" + oldItem.IngredientsPortuguese + "_ to <BR>_" + newItem.IngredientsPortuguese + "_<BR>"
	}
	if oldItem.IngredientsSpanish != newItem.IngredientsSpanish {
		differencesString += "**IngredientsSpanish** changed from <BR>_" + oldItem.IngredientsSpanish + "_ to <BR>_" + newItem.IngredientsSpanish + "_<BR>"
	}
	if oldItem.IngredientsScandinavian != newItem.IngredientsScandinavian {
		differencesString += "**IngredientsScandinavian** changed from <BR>_" + oldItem.IngredientsScandinavian + "_ to <BR>_" + newItem.IngredientsScandinavian + "_<BR>"
	}

	return differencesString
}
