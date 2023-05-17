package utils

import (
	"fmt"
	"list_of_ingredients_producer/sap_api_wrapper"
	"list_of_ingredients_producer/teams_notifier"
)

type ItemCodeAndQuantity struct {
	ItemCode string
	Quantity float64
}

type BillOfMaterials []sap_api_wrapper.SapApiBillOfMaterialEntry

// Retrieves all the needed items from SAP and updates the list of ingredients for each item
func HandleAllItemsListOfIngredients() error {
	//items, err := GetItemDataFromSap("0022030034-1")
	items, err := GetItemDataFromSap("0021030129")
	if err != nil {
		return fmt.Errorf("error getting item data from SAP: %v", err)
	}

	for _, item := range items.Value {
		itemError := func() error {

			// Create a sales item so we can store the raw materials within it
			salesItem := item

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
			sortedSliceOfMaterials := SortMaterialsByQuantity(rawMaterialsMap)

			salesItem, err := HandleConcatAllListOfIngredients(sortedSliceOfMaterials, salesItem, totalQuantity)
			if err != nil {
				return err
			}

			err = sap_api_wrapper.SetItemData(salesItem)
			if err != nil {
				return err
			}

			return nil
		}()

		if itemError != nil {
			teams_notifier.SendUnknownErrorToTeams(itemError)
		}
	}

	return nil
}
