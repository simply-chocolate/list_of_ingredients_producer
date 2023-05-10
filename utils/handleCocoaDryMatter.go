package utils

import (
	"strings"
)

func getCocoaDryMatterString(materialsOnProduct []RawMaterial, allRawMaterials map[string]map[string]string, languageCode string) string {

	cocoaDryMatterString := ""
	hasDarkChocolate := false
	hasMilkChocolate := false

	for _, material := range materialsOnProduct {

		materialInfo, exists := allRawMaterials[material.ItemCode]
		if !exists {
			panic("error: could not find raw material in allRawMaterials")
		}

		//fmt.Println("ItemName: ", materialInfo["ItemName"])
		if strings.Contains(materialInfo["ItemName"], "chokolade") || strings.Contains(materialInfo["ItemName"], "Chokolade") {
			if strings.Contains(materialInfo["ItemName"], "Mørk") || strings.Contains(materialInfo["ItemName"], "mørk") {
				hasDarkChocolate = true
			}
			if strings.Contains(materialInfo["ItemName"], "Mælk") || strings.Contains(materialInfo["ItemName"], "mælk") {
				hasMilkChocolate = true
			}
		}
	}

	darkChocolateCocoaDryMatterMap := map[string]string{
		"DA_SE_NO": "Mindst 60% kakaotørstofindhold/kakaoinnehåll i den mørke/mörka chokolade/choklad/sjokolade. ",
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
		"DA_SE_NO": "Mindst 35% kakaotørstofindhold/kakaoinnehåll i mælke/mjölkchokoladen/-choklad/-sjokoladen. ",
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
	if hasMilkChocolate {
		cocoaDryMatterString += milkChocolateCocoaDryMatterMap[languageCode]
	}

	return cocoaDryMatterString
}
