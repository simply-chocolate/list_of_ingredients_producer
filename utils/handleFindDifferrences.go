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
	if oldItem.ContainmentLevelAlmonds != newItem.ContainmentLevelAlmonds {
		differencesString += "**ContainmentLevelAlmonds** changed from _" + oldItem.ContainmentLevelAlmonds + "_ to _" + newItem.ContainmentLevelAlmonds + "_<BR>"
	}
	if oldItem.ContainmentLevelBrazilNut != newItem.ContainmentLevelBrazilNut {
		differencesString += "**ContainmentLevelBrazilNut** changed from _" + oldItem.ContainmentLevelBrazilNut + "_ to _" + newItem.ContainmentLevelBrazilNut + "_<BR>"
	}
	if oldItem.ContainmentLevelCashew != newItem.ContainmentLevelCashew {
		differencesString += "**ContainmentLevelCashew** changed from _" + oldItem.ContainmentLevelCashew + "_ to _" + newItem.ContainmentLevelCashew + "_<BR>"
	}
	if oldItem.ContainmentLevelCelery != newItem.ContainmentLevelCelery {
		differencesString += "**ContainmentLevelCelery** changed from _" + oldItem.ContainmentLevelCelery + "_ to _" + newItem.ContainmentLevelCelery + "_<BR>"
	}
	if oldItem.ContainmentLevelCrustacea != newItem.ContainmentLevelCrustacea {
		differencesString += "**ContainmentLevelCrustacea** changed from _" + oldItem.ContainmentLevelCrustacea + "_ to _" + newItem.ContainmentLevelCrustacea + "_<BR>"
	}
	if oldItem.ContainmentLevelEgg != newItem.ContainmentLevelEgg {
		differencesString += "**ContainmentLevelEgg** changed from _" + oldItem.ContainmentLevelEgg + "_ to _" + newItem.ContainmentLevelEgg + "_<BR>"
	}
	if oldItem.ContainmentLevelFish != newItem.ContainmentLevelFish {
		differencesString += "**ContainmentLevelFish** changed from _" + oldItem.ContainmentLevelFish + "_ to _" + newItem.ContainmentLevelFish + "_<BR>"
	}
	if oldItem.ContainmentLevelGluten != newItem.ContainmentLevelGluten {
		differencesString += "**ContainmentLevelGluten** changed from _" + oldItem.ContainmentLevelGluten + "_ to _" + newItem.ContainmentLevelGluten + "_<BR>"
	}
	if oldItem.ContainmentLevelHazelnut != newItem.ContainmentLevelHazelnut {
		differencesString += "**ContainmentLevelHazelnut** changed from _" + oldItem.ContainmentLevelHazelnut + "_ to _" + newItem.ContainmentLevelHazelnut + "_<BR>"
	}
	if oldItem.ContainmentLevelLupine != newItem.ContainmentLevelLupine {
		differencesString += "**ContainmentLevelLupine** changed from _" + oldItem.ContainmentLevelLupine + "_ to _" + newItem.ContainmentLevelLupine + "_<BR>"
	}
	if oldItem.ContainmentLevelMilk != newItem.ContainmentLevelMilk {
		differencesString += "**ContainmentLevelMilk** changed from _" + oldItem.ContainmentLevelMilk + "_ to _" + newItem.ContainmentLevelMilk + "_<BR>"
	}
	if oldItem.ContainmentLevelMollusks != newItem.ContainmentLevelMollusks {
		differencesString += "**ContainmentLevelMollusks** changed from _" + oldItem.ContainmentLevelMollusks + "_ to _" + newItem.ContainmentLevelMollusks + "_<BR>"
	}
	if oldItem.ContainmentLevelMustard != newItem.ContainmentLevelMustard {
		differencesString += "**ContainmentLevelMustard** changed from _" + oldItem.ContainmentLevelMustard + "_ to _" + newItem.ContainmentLevelMustard + "_<BR>"
	}
	if oldItem.ContainmentLevelPeanut != newItem.ContainmentLevelPeanut {
		differencesString += "**ContainmentLevelPeanut** changed from _" + oldItem.ContainmentLevelPeanut + "_ to _" + newItem.ContainmentLevelPeanut + "_<BR>"
	}
	if oldItem.ContainmentLevelPecan != newItem.ContainmentLevelPecan {
		differencesString += "**ContainmentLevelPecan** changed from _" + oldItem.ContainmentLevelPecan + "_ to _" + newItem.ContainmentLevelPecan + "_<BR>"
	}
	if oldItem.ContainmentLevelPistachio != newItem.ContainmentLevelPistachio {
		differencesString += "**ContainmentLevelPistachio** changed from _" + oldItem.ContainmentLevelPistachio + "_ to _" + newItem.ContainmentLevelPistachio + "_<BR>"
	}
	if oldItem.ContainmentLevelQueenslandNut != newItem.ContainmentLevelQueenslandNut {
		differencesString += "**ContainmentLevelQueenslandNut** changed from _" + oldItem.ContainmentLevelQueenslandNut + "_ to _" + newItem.ContainmentLevelQueenslandNut + "_<BR>"
	}
	if oldItem.ContainmentLevelSesameSeeds != newItem.ContainmentLevelSesameSeeds {
		differencesString += "**ContainmentLevelSesameSeeds** changed from _" + oldItem.ContainmentLevelSesameSeeds + "_ to _" + newItem.ContainmentLevelSesameSeeds + "_<BR>"
	}
	if oldItem.ContainmentLevelSoy != newItem.ContainmentLevelSoy {
		differencesString += "**ContainmentLevelSoy** changed from _" + oldItem.ContainmentLevelSoy + "_ to _" + newItem.ContainmentLevelSoy + "_<BR>"
	}
	if oldItem.ContainmentLevelSulfurDioxideAndSulfites != newItem.ContainmentLevelSulfurDioxideAndSulfites {
		differencesString += "**ContainmentLevelSulfurDioxideAndSulfites** changed from _" + oldItem.ContainmentLevelSulfurDioxideAndSulfites + "_ to _" + newItem.ContainmentLevelSulfurDioxideAndSulfites + "_<BR>"
	}
	if oldItem.ContainmentLevelWalnut != newItem.ContainmentLevelWalnut {
		differencesString += "**ContainmentLevelWalnut** changed from _" + oldItem.ContainmentLevelWalnut + "_ to _" + newItem.ContainmentLevelWalnut + "_<BR>"
	}
	if oldItem.CowFree != newItem.CowFree {
		differencesString += "**CowFree** changed from _" + oldItem.CowFree + "_ to _" + newItem.CowFree + "_<BR>"
	}
	if oldItem.GlutenFree != newItem.GlutenFree {
		differencesString += "**GlutenFree** changed from _" + oldItem.GlutenFree + "_ to _" + newItem.GlutenFree + "_<BR>"
	}
	if oldItem.GMOFree != newItem.GMOFree {
		differencesString += "**GMOFree** changed from _" + oldItem.GMOFree + "_ to _" + newItem.GMOFree + "_<BR>"
	}
	if oldItem.LactoseFree != newItem.LactoseFree {
		differencesString += "**LactoseFree** changed from _" + oldItem.LactoseFree + "_ to _" + newItem.LactoseFree + "_<BR>"
	}
	if oldItem.PigFree != newItem.PigFree {
		differencesString += "**PigFree** changed from _" + oldItem.PigFree + "_ to _" + newItem.PigFree + "_<BR>"
	}
	if oldItem.Vegan != newItem.Vegan {
		differencesString += "**Vegan** changed from _" + oldItem.Vegan + "_ to _" + newItem.Vegan + "_<BR>"
	}
	if oldItem.Vegetarian != newItem.Vegetarian {
		differencesString += "**Vegetarian** changed from _" + oldItem.Vegetarian + "_ to _" + newItem.Vegetarian + "_<BR>"
	}

	return differencesString
}
