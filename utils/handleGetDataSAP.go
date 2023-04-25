package utils

import (
	"fmt"
	"list_of_ingredients_producer/sap_api_wrapper"
)

func GetItemDataFromSap() (sap_api_wrapper.SapApiGetItemsDataResults, error) {
	resp, err := sap_api_wrapper.SapApiGetItemsData_AllPages(sap_api_wrapper.SapApiQueryParams{
		Filter: "ItemCode eq '0021050008'",
	})
	if err != nil {
		return sap_api_wrapper.SapApiGetItemsDataResults{}, err
	}

	return *resp.Body, nil
}

func GetMixCaseItemsFromSap(itemCode string) (sap_api_wrapper.SapApiGetMixCaseDataResult, error) {
	resp, err := sap_api_wrapper.SapApiGetMixCaseData_AllPages(sap_api_wrapper.SapApiQueryParams{
		FatherItemCode: fmt.Sprintf("'%s'", itemCode),
	})
	if err != nil {
		return sap_api_wrapper.SapApiGetMixCaseDataResult{}, err
	}

	return *resp.Body, nil
}
