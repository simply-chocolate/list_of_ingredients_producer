package utils

import (
	"fmt"
	"list_of_ingredients_producer/sap_api_wrapper"
)

func GetItemsDataFromSap() (sap_api_wrapper.SapApiGetItemsDataResults, error) {
	resp, err := sap_api_wrapper.SapApiGetItemsData_AllPages(sap_api_wrapper.SapApiQueryParams{
		Filter: "U_CCF_Update_NUT eq 'Y'",
	})
	if err != nil {
		return sap_api_wrapper.SapApiGetItemsDataResults{}, err
	}

	return *resp.Body, nil
}

func GetRawMaterialsItemDataFromSap() (map[string]map[string]string, error) {
	resp, err := sap_api_wrapper.SapApiGetItemsData_AllPages(sap_api_wrapper.SapApiQueryParams{
		Filter: "U_CCF_Type eq 'Råvare'",
	})
	if err != nil {
		return map[string]map[string]string{}, err
	}

	// Make resp.body.value into a map
	// This is so we can easily check if an itemcode is in the list

	rawMaterialItemDataMap := make(map[string]map[string]string)
	for _, item := range resp.Body.Value {
		rawMaterialItemData := make(map[string]string)

		rawMaterialItemData["ItemName"] = item.ItemName

		// Lists of ingredients
		rawMaterialItemData["DA_SE_NO"] = item.IngredientsScandinavian
		rawMaterialItemData["DA"] = item.IngredientsDanish
		rawMaterialItemData["FI"] = item.IngredientsFinnish
		rawMaterialItemData["EN"] = item.IngredientsEnglish
		rawMaterialItemData["DE"] = item.IngredientsGerman
		rawMaterialItemData["NL"] = item.IngredientsDutch
		rawMaterialItemData["FR"] = item.IngredientsFrench
		rawMaterialItemData["PT"] = item.IngredientsPortuguese
		rawMaterialItemData["IT"] = item.IngredientsItalian
		rawMaterialItemData["ES"] = item.IngredientsSpanish

		// Allergen containment information
		rawMaterialItemData["Gluten"] = item.ContainmentLevelGluten
		rawMaterialItemData["Crustacea"] = item.ContainmentLevelCrustacea
		rawMaterialItemData["Egg"] = item.ContainmentLevelEgg
		rawMaterialItemData["Fish"] = item.ContainmentLevelFish
		rawMaterialItemData["Peanut"] = item.ContainmentLevelPeanut
		rawMaterialItemData["Soy"] = item.ContainmentLevelSoy
		rawMaterialItemData["Milk"] = item.ContainmentLevelMilk
		rawMaterialItemData["Almonds"] = item.ContainmentLevelAlmonds
		rawMaterialItemData["Hazelnut"] = item.ContainmentLevelHazelnut
		rawMaterialItemData["Walnut"] = item.ContainmentLevelWalnut
		rawMaterialItemData["Cashew"] = item.ContainmentLevelCashew
		rawMaterialItemData["Pecan"] = item.ContainmentLevelPecan
		rawMaterialItemData["BrazilNut"] = item.ContainmentLevelBrazilNut
		rawMaterialItemData["Pistachio"] = item.ContainmentLevelPistachio
		rawMaterialItemData["QueenslandNut"] = item.ContainmentLevelQueenslandNut
		rawMaterialItemData["Celery"] = item.ContainmentLevelCelery

		// Nutritional information
		rawMaterialItemData["EnergyKJ"] = item.EnergyInkJ
		rawMaterialItemData["EnergyKcal"] = item.EnergyInKcal
		rawMaterialItemData["Fat"] = item.NutritionalFatValue
		rawMaterialItemData["FattyAcids"] = item.NutritionalFattyAcidsValue
		rawMaterialItemData["CarboHydrates"] = item.NutritionalCarboHydratesValue
		rawMaterialItemData["Sugar"] = item.NutritionalSugarValue
		rawMaterialItemData["Protein"] = item.NutritionalProteinValue
		rawMaterialItemData["Salt"] = item.NutritionalSaltValue

		rawMaterialItemDataMap[item.ItemCode] = rawMaterialItemData

	}

	return rawMaterialItemDataMap, nil
}

func GetItemDataFromSap(itemCode string) (sap_api_wrapper.SapApiGetItemsDataResults, error) {
	resp, err := sap_api_wrapper.SapApiGetItemsData_AllPages(sap_api_wrapper.SapApiQueryParams{
		Filter: fmt.Sprintf("ItemCode eq '%v'", itemCode),
	})
	if err != nil {
		return sap_api_wrapper.SapApiGetItemsDataResults{}, err
	}

	return *resp.Body, nil
}

func GetBillOfMaterialFromSap(itemCode string) (sap_api_wrapper.SapApiGetBillOfMaterialsDataResult, error) {
	resp, err := sap_api_wrapper.SapApiGetBillOfMaterialsData_AllPages(sap_api_wrapper.SapApiQueryParams{
		FatherItemCode: fmt.Sprintf("'%s'", itemCode),
	})
	if err != nil {
		return sap_api_wrapper.SapApiGetBillOfMaterialsDataResult{}, err
	}

	return *resp.Body, nil
}
