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

// TODO: Vi skal kunne gå et niveau ned i styklisterne for at finde alle de sammensatte råvarers underråvarer.
// Til det skal vi have oprettet alle de råvarer, som er en del af en sammensat ingrediens. fx "sukker, soya lecitin".
// Nogle af disse findes allerede i SAP, hvortil vi blot skal oprette en stykliste på de sammensatte råvarer.
// For at kunne overholde lov om anprisninger skal vi finde ud af hvordan vi laver en lille "lækker" tekst oppe i toppen,
// hvor vi kan angive hvor mange % af hver sammensat råvare der er i produktet.
// For at kunne komme i mål med dette, kræver det at vi kan skrive navnet på de ingrediens, på alle de sprog vi bruger.
// U_CCF_IngredientName_DA_SE_NO, U_CCF_IngredientName_FI, U_CCF_IngredientName_EN, U_CCF_IngredientName_DE, U_CCF_IngredientName_NL, U_CCF_IngredientName_FR, U_CCF_IngredientName_PT, U_CCF_IngredientName_IT, U_CCF_IngredientName_ES
// For mørk chokolade skal man så bare skrive "Mørk/Mörk chokolade/choklad/sjokolade" i U_CCF_IngredientName_DA_SE_NO.
// Når du så bruger ingrediensen i et produkt, vil der i toppen stå "Mørk/Mörk chokolade/choklad/sjokolade (xx%) med Kastajnjer (xx%) og Sukker (xx%)".

// Retrieves all the needed items from SAP and updates the list of ingredients for each item
func HandleAllItemsListOfIngredients() error {
	items, err := GetItemsDataFromSap()
	//items, err := GetItemDataFromSap("0031030126")
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

			differences := FindDifferences(item, salesItem)

			if len(differences) > 0 {
				teams_notifier.SendProductUpdatedMessageToTemas(salesItem.ItemCode, FindDifferences(item, salesItem))
				err = sap_api_wrapper.SetItemData(salesItem)
				if err != nil {
					return err
				}
			}

			return nil
		}()

		if itemError != nil {
			teams_notifier.SendUnknownErrorToTeams(itemError)
		}
	}

	return nil
}
