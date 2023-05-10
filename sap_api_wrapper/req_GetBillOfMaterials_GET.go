package sap_api_wrapper

import (
	"fmt"
	"time"
)

type SapApiGetBillOfMaterialsDataResult struct {
	Value    []SapApiBillOfMaterialEntry `json:"value"`
	NextLink string                      `json:"odata.nextLink"`
}

type SapApiBillOfMaterialEntry struct {
	ItemCode string  `json:"ItemCode"`
	Quantity float64 `json:"Quantity"`
	Type     string  `json:"U_CCF_Type"`
}

type SapApiGetBillOfMaterialsDataReturn struct {
	Body *SapApiGetBillOfMaterialsDataResult
}

func SapApiGetBillOfMaterialsData(params SapApiQueryParams) (SapApiGetBillOfMaterialsDataReturn, error) {
	for i := 0; i < 200; i++ {
		client, err := GetSapApiAuthClient()
		if err != nil {
			fmt.Println("Error getting an authenticaed client")
			return SapApiGetBillOfMaterialsDataReturn{}, err
		}

		resp, err := client.
			//DevMode().
			R().
			SetResult(SapApiGetBillOfMaterialsDataResult{}).
			SetQueryParams(params.AsReqParams()).
			Get("SQLQueries('CQ10003')/List")
		if err != nil {
			fmt.Println(err)
			return SapApiGetBillOfMaterialsDataReturn{}, err
		}

		if resp.IsError() {
			if resp.StatusCode != 403 {
				//fmt.Printf("Dumping SAP Error %v\n", resp.Dump())
				return SapApiGetBillOfMaterialsDataReturn{}, fmt.Errorf("error getting BillOfMaterialsContent from to sap. unexpected errorcode. StatusCode :%v Status: %v", resp.StatusCode, resp.Status)
			} else {
				time.Sleep(100 * time.Millisecond)
			}

		} else {
			return SapApiGetBillOfMaterialsDataReturn{
				Body: resp.Result().(*SapApiGetBillOfMaterialsDataResult),
			}, nil
		}
	}
	return SapApiGetBillOfMaterialsDataReturn{}, fmt.Errorf("error getting invoices from SAP. Tried 200 times and couldn't get through")
}

func SapApiGetBillOfMaterialsData_AllPages(params SapApiQueryParams) (SapApiGetBillOfMaterialsDataReturn, error) {
	res := SapApiGetBillOfMaterialsDataResult{}
	for page := 0; ; page++ {
		params.Skip = page * 20

		getItemsRes, err := SapApiGetBillOfMaterialsData(params)
		if err != nil {
			return SapApiGetBillOfMaterialsDataReturn{}, err
		}

		res.Value = append(res.Value, getItemsRes.Body.Value...)

		if getItemsRes.Body.NextLink == "" {
			break
		}
	}

	return SapApiGetBillOfMaterialsDataReturn{
		Body: &res,
	}, nil
}
