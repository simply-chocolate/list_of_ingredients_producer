package utils

import "list_of_ingredients_producer/sap_api_wrapper"

func FindDifferences(oldItem sap_api_wrapper.SapApiItemsData, newItem sap_api_wrapper.SapApiItemsData) string {
	differencesString := ""

	if oldItem.EnergyInKcal != newItem.EnergyInKcal {
		differencesString += "**EnergyInKcal** changed from " + oldItem.EnergyInKcal + " to " + newItem.EnergyInKcal + "<BR>"
	}
	if oldItem.EnergyInkJ != newItem.EnergyInkJ {
		differencesString += "**EnergyInkJ** changed from " + oldItem.EnergyInkJ + " to " + newItem.EnergyInkJ + "<BR>"
	}
	if oldItem.NutritionalFatValue != newItem.NutritionalFatValue {
		differencesString += "**Fat** changed from " + oldItem.NutritionalFatValue + " to " + newItem.NutritionalFatValue + "<BR>"
	}
	if oldItem.NutritionalFattyAcidsValue != newItem.NutritionalFattyAcidsValue {
		differencesString += "**FattyAcids** changed from " + oldItem.NutritionalFattyAcidsValue + " to " + newItem.NutritionalFattyAcidsValue + "<BR>"
	}
	if oldItem.NutritionalCarboHydratesValue != newItem.NutritionalCarboHydratesValue {
		differencesString += "**CarboHydrates** changed from " + oldItem.NutritionalCarboHydratesValue + " to " + newItem.NutritionalCarboHydratesValue + "<BR>"
	}
	if oldItem.NutritionalSugarValue != newItem.NutritionalSugarValue {
		differencesString += "**Sugar** changed from " + oldItem.NutritionalSugarValue + " to " + newItem.NutritionalSugarValue + "<BR>"
	}
	if oldItem.NutritionalProteinValue != newItem.NutritionalProteinValue {
		differencesString += "**Protein** changed from " + oldItem.NutritionalProteinValue + " to " + newItem.NutritionalProteinValue + "<BR>"
	}
	if oldItem.NutritionalSaltValue != newItem.NutritionalSaltValue {
		differencesString += "**Salt** changed from " + oldItem.NutritionalSaltValue + " to " + newItem.NutritionalSaltValue + "<BR>"
	}
	if oldItem.IngredientsDanish != newItem.IngredientsDanish {
		differencesString += "**IngredientsDanish** changed from " + oldItem.IngredientsDanish + " to " + newItem.IngredientsDanish + "<BR>"
	}
	if oldItem.IngredientsDutch != newItem.IngredientsDutch {
		differencesString += "**IngredientsDutch** changed from " + oldItem.IngredientsDutch + " to " + newItem.IngredientsDutch + "<BR>"
	}
	if oldItem.IngredientsEnglish != newItem.IngredientsEnglish {
		differencesString += "**IngredientsEnglish** changed from " + oldItem.IngredientsEnglish + " to " + newItem.IngredientsEnglish + "<BR>"
	}
	if oldItem.IngredientsFinnish != newItem.IngredientsFinnish {
		differencesString += "**IngredientsFinnish** changed from " + oldItem.IngredientsFinnish + " to " + newItem.IngredientsFinnish + "<BR>"
	}
	if oldItem.IngredientsFrench != newItem.IngredientsFrench {
		differencesString += "**IngredientsFrench** changed from " + oldItem.IngredientsFrench + " to " + newItem.IngredientsFrench + "<BR>"
	}
	if oldItem.IngredientsGerman != newItem.IngredientsGerman {
		differencesString += "**IngredientsGerman** changed from " + oldItem.IngredientsGerman + " to " + newItem.IngredientsGerman + "<BR>"
	}
	if oldItem.IngredientsItalian != newItem.IngredientsItalian {
		differencesString += "**IngredientsItalian** changed from " + oldItem.IngredientsItalian + " to " + newItem.IngredientsItalian + "<BR>"
	}
	if oldItem.IngredientsPortuguese != newItem.IngredientsPortuguese {
		differencesString += "**IngredientsPortuguese** changed from " + oldItem.IngredientsPortuguese + " to " + newItem.IngredientsPortuguese + "<BR>"
	}
	if oldItem.IngredientsSpanish != newItem.IngredientsSpanish {
		differencesString += "**IngredientsSpanish** changed from " + oldItem.IngredientsSpanish + " to " + newItem.IngredientsSpanish + "<BR>"
	}
	if oldItem.IngredientsScandinavian != newItem.IngredientsScandinavian {
		differencesString += "**IngredientsScandinavian** changed from " + oldItem.IngredientsScandinavian + " to " + newItem.IngredientsScandinavian + "<BR>"
	}

	return differencesString
}
