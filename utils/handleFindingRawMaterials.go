package utils

import (
	"list_of_ingredients_producer/sap_api_wrapper"
)

// In here we need to find the raw materials for the recipe

// We need to take an argument that is the itemcode
// Then we will use the itemcode to find the bill of materials
// Then we will check the bill of materials to see if it is a raw material
// If the typing is raw material, we will add it to the list of raw materials for the itemcode
//
//	We are going to store the list of ingredients for the raw material in a map containing each language
//
// If the typing is not raw material, we will call the function again with the new itemcode
// We will need to return a list of raw materials for the itemcode
type RawMaterial struct {
	ItemCode                string `json:"ItemCode"`
	FatherQuantity          float64
	Quantity                float64 `json:"Quantity"`
	IngredientsScandinavian string  `json:"U_CCF_Ingrediens_DA_SE_NO"`
	IngredientsFinnish      string  `json:"U_CCF_Ingrediens_FI"`
	IngredientsEnglish      string  `json:"U_CCF_Ingrediens_EN"`
	IngredientsGerman       string  `json:"U_CCF_Ingrediens_DE"`
	IngredientsDutch        string  `json:"U_CCF_Ingrediens_NL"`
	IngredientsFrench       string  `json:"U_CCF_Ingrediens_FR"`
	IngredientsPortuguese   string  `json:"U_CCF_Ingrediens_PT"`
	IngredientsItalian      string  `json:"U_CCF_Ingrediens_IT"`
	IngredientsSpanish      string  `json:"U_CCF_Ingrediens_ES"`
}

type ItemCodeAndQuantity struct {
	ItemCode string
	Quantity float64
}

type BillOfMaterials []sap_api_wrapper.SapApiBillOfMaterialEntry

type SalesItem struct {
	ItemCode                string
	IngredientsScandinavian string `json:"U_CCF_Ingrediens_DA_SE_NO"`
	IngredientsFinnish      string `json:"U_CCF_Ingrediens_FI"`
	IngredientsEnglish      string `json:"U_CCF_Ingrediens_EN"`
	IngredientsGerman       string `json:"U_CCF_Ingrediens_DE"`
	IngredientsDutch        string `json:"U_CCF_Ingrediens_NL"`
	IngredientsFrench       string `json:"U_CCF_Ingrediens_FR"`
	IngredientsPortuguese   string `json:"U_CCF_Ingrediens_PT"`
	IngredientsItalian      string `json:"U_CCF_Ingrediens_IT"`
	IngredientsSpanish      string `json:"U_CCF_Ingrediens_ES"`
}

func GetAllBillOfMaterials() {
	items, err := GetItemDataFromSap("0022030034-1")
	//items, err := GetItemDataFromSap("0021050008")
	if err != nil {
		panic(err)
	}

	for _, item := range items.Value {
		// Create a sales item so we can store the raw materials within it
		var salesItem SalesItem
		salesItem.ItemCode = item.ItemCode

		var billOfMaterials BillOfMaterials
		var starterItem ItemCodeAndQuantity
		starterItem.ItemCode = item.ItemCode
		starterItem.Quantity = 1.0
		listOfItemCodesToCheck := []ItemCodeAndQuantity{starterItem}

		// Have a list of itemcodes that we need to get BillOfMaterials for and for each call, we append all the HF's to that list
		for i := 0; len(listOfItemCodesToCheck) > 0; {
			var HFs []ItemCodeAndQuantity

			billOfMaterials, HFs, err = GetBillOfMaterials(listOfItemCodesToCheck[i].ItemCode, billOfMaterials, listOfItemCodesToCheck[i].Quantity)
			if err != nil {
				panic(err)
			}

			// Remove the itemcode from the list of itemcodes to check
			listOfItemCodesToCheck = listOfItemCodesToCheck[1:]

			// Add any new itemcodes to the list of itemcodes to check
			listOfItemCodesToCheck = append(listOfItemCodesToCheck, HFs...)
		}

		rawMaterialsMap, totalQuantity := MapRawMaterials(billOfMaterials)

		// Sort the map by quantity
		sortedSliceOfMaterials := SortMaterialsByQuantity(rawMaterialsMap)

		HandleConcatListOfIngredients(sortedSliceOfMaterials, salesItem, totalQuantity)

	}
}
