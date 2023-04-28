package utils

import (
	"fmt"
	"list_of_ingredients_producer/sap_api_wrapper"
)

func GetItemsDataFromSap() (sap_api_wrapper.SapApiGetItemsDataResults, error) {
	resp, err := sap_api_wrapper.SapApiGetItemsData_AllPages(sap_api_wrapper.SapApiQueryParams{
		Select: []string{"ItemCode", "ItemName"},
		Filter: "U_CCF_Type eq 'Færdigvare'",
	})
	if err != nil {
		return sap_api_wrapper.SapApiGetItemsDataResults{}, err
	}

	return *resp.Body, nil
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
