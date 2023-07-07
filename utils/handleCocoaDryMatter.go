package utils

import (
	"strings"
)

func getCocoaDryMatterString(materialsOnProduct []RawMaterial, allRawMaterials map[string]map[string]string, languageCode string) string {

	cocoaDryMatterString := ""
	hasDarkChocolate := false
	hasVeryDarkChocolate := false
	hasMilkChocolate := false

	for _, material := range materialsOnProduct {

		materialInfo, exists := allRawMaterials[material.ItemCode]
		if !exists {
			panic("error: could not find raw material in allRawMaterials")
		}

		// TODO: Handle this from SAP. We need a field on the raw material that tells us if it is 60%, 70%, 35% or 100% cocoa dry matter.
		// Also weather its "dark", "milk" or "cocoa nibs"
		if strings.Contains(materialInfo["ItemName"], "chokolade") || strings.Contains(materialInfo["ItemName"], "Chokolade") {
			if strings.Contains(materialInfo["ItemName"], "Mørk") || strings.Contains(materialInfo["ItemName"], "mørk") {
				if strings.Contains(materialInfo["ItemName"], "70%") {
					hasVeryDarkChocolate = true
				} else {
					hasDarkChocolate = true
				}
			}
			if strings.Contains(materialInfo["ItemName"], "Mælk") || strings.Contains(materialInfo["ItemName"], "mælk") {
				hasMilkChocolate = true
			}
		}
	}

	veryDarkChocolateCocoaDryMatterMap := map[string]string{
		"DA_SE_NO": "Mørk chokolade: Mindst 70% kakaotørstof. ",
		"DA":       "Mørk chokolade: Mindst 70% kakaotørstof. ",
		"EN":       "Min. 70% dry cocoa solids in the dark chocolate. ",
		"FI":       "Tumman suklaan kaakaokuiva-aineen pitoisuus vähintään 70%. ",
		"DE":       "Mindestens 70% Kakaotrockenmasse in dunkler Schokolade. ",
		"NL":       "Ten minste 70% droge cacaobestanddelen in de pure chocola. ",
		"FR":       "Minimum 70% de cacao sec dans le chocolat noir. ",
		"PT":       "Pelo menos 70% de teor de cacau no chocolate escuro. ",
		"IT":       "Contenuto di cacao minimo 70% della sostanza secca nel cioccolato fondente. ",
		"ES":       "Al menos 70% de sólidos secos de cacao en el chocolate amargo. ",
	}

	darkChocolateCocoaDryMatterMap := map[string]string{
		"DA_SE_NO": "Mørk chokolade: Mindst 60% kakaotørstof. ",
		"DA":       "Mørk chokolade: Mindst 60% kakaotørstof. ",
		"EN":       "Min. 60% dry cocoa solids in the dark chocolate. ",
		"FI":       "Tumman suklaan kaakaokuiva-aineen pitoisuus vähintään 60%. ",
		"DE":       "Mindestens 60% Kakaotrockenmasse in dunkler Schokolade. ",
		"NL":       "Ten minste 60% droge cacaobestanddelen in de pure chocola. ",
		"FR":       "Minimum 60% de cacao sec dans le chocolat noir. ",
		"PT":       "Pelo menos 60% de teor de cacau no chocolate escuro. ",
		"IT":       "Contenuto di cacao minimo 60% della sostanza secca nel cioccolato fondente. ",
		"ES":       "Al menos 60% de sólidos secos de cacao en el chocolate amargo. ",
	}
	milkChocolateCocoaDryMatterMap := map[string]string{
		"DA_SE_NO": "Mælke-/mjölk-chokolade: Mindst 35% kakaotørstof. ",
		"DA":       "Mælke-/mjölk-chokolade: Mindst 35% kakaotørstof. ",
		"EN":       "Min. 35% dry cocoa solids in the milk chocolate. ",
		"FI":       "Maitosuklaan kaakaokuiva-aineen pitoisuus vähintään 35%. ",
		"DE":       "Mindestens 35% Kakaotrockenmasse in Milchschokolade. ",
		"NL":       "Ten minste 35% droge cacaobestanddelen in de melkchocola. ",
		"FR":       "Minimum 35% de cacao sec dans le chocolat au lait. ",
		"PT":       "Pelo menos 35% de teor de cacau no chocolate de leite. ",
		"IT":       "Contenuto di cacao minimo 35% della sostanza secca nel cioccolato al latte. ",
		"ES":       "Al menos 35% de sólidos secos de cacao en el chocolate con leche. ",
	}

	if hasDarkChocolate {
		cocoaDryMatterString += darkChocolateCocoaDryMatterMap[languageCode]
	}
	if hasVeryDarkChocolate {
		cocoaDryMatterString += veryDarkChocolateCocoaDryMatterMap[languageCode]
	}
	if hasMilkChocolate {
		cocoaDryMatterString += milkChocolateCocoaDryMatterMap[languageCode]
	}

	return cocoaDryMatterString
}
