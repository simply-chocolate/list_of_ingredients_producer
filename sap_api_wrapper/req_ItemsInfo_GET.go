package sap_api_wrapper

import (
	"fmt"
)

type SapApiGetItemsDataResults struct {
	Value    []SapApiItemsData `json:"value"`
	NextLink string            `json:"odata.nextLink"`
}

/*
type ItemUnitOfMeasurement struct {
	UoMType    string `json:"UoMType"`
	UoMEntry   int    `json:"UoMEntry"`
	Length     int    `json:"Length1"`
	LengthUnit int    `json:"Length1Unit"` // For dimensions [1 = mm?, 2 = cm, 3 = m?]
	Width      int    `json:"Width1"`
	WidthUnit  int    `json:"Width1Unit"` // For dimensions [1 = mm?, 2 = cm, 3 = m?]
	Height     int    `json:"Height1"`
	HeightUnit int    `json:"Height1Unit"` // For dimensions [1 = mm?, 2 = cm, 3 = m?]
	Weight     int    `json:"Weight1"`
	WeightUnit int    `json:"Weight1Unit"` // For weight [1 = ?, 2 = g?, 3 = kg]
	// Maybe the units is 1 list, so 2 = cm and 3 = kg no matter what type of input?
}
*/

type SapApiItemsData struct {
	TypeOfProduct string `json:"U_CCF_Type"` // If this is Equal to "Kampagne" then it should not have a BaseUnit ItemCode
	ItemCode      string `json:"ItemCode"`
	ItemName      string `json:"ItemName"`

	// Lists of ingredients
	IngredientsScandinavian string `json:"U_CCF_Ingrediens_DA_SE_NO"`
	IngredientsFinnish      string `json:"U_CCF_Ingrediens_FI"`
	IngredientsEnglish      string `json:"U_CCF_Ingrediens_EN"`
	IngredientsGerman       string `json:"U_CCF_Ingrediens_DE"`
	IngredientsDutch        string `json:"U_CCF_Ingrediens_NL"`
	IngredientsFrench       string `json:"U_CCF_Ingrediens_FR"`
	IngredientsPortuguese   string `json:"U_CCF_Ingrediens_PT"`
	IngredientsItalian      string `json:"U_CCF_Ingrediens_IT"`
	IngredientsSpanish      string `json:"U_CCF_Ingrediens_ES"`

	// Allergen containment information
	ContainmentLevelGluten                   string `json:"U_BOYX_gluten"`
	ContainmentLevelCrustacea                string `json:"U_BOYX_Krebsdyr"`
	ContainmentLevelEgg                      string `json:"U_BOYX_aag"`
	ContainmentLevelFish                     string `json:"U_BOYX_fisk"`
	ContainmentLevelPeanut                   string `json:"U_BOYX_JN"`
	ContainmentLevelSoy                      string `json:"U_BOYX_soja"`
	ContainmentLevelMilk                     string `json:"U_BOYX_ML"`
	ContainmentLevelAlmonds                  string `json:"U_BOYX_mandel"`
	ContainmentLevelHazelnut                 string `json:"U_BOYX_hassel"`
	ContainmentLevelWalnut                   string `json:"U_BOYX_val"`
	ContainmentLevelCashew                   string `json:"U_BOYX_Cashe"`
	ContainmentLevelPecan                    string `json:"U_BOYX_Pekan"`
	ContainmentLevelBrazilNut                string `json:"U_BOYX_peka"`
	ContainmentLevelPistachio                string `json:"U_BOYX_Pistacie"`
	ContainmentLevelQueenslandNut            string `json:"U_BOYX_Queensland"`
	ContainmentLevelCelery                   string `json:"U_BOYX_Selleri"`
	ContainmentLevelMustard                  string `json:"U_BOYX_Sennep"`
	ContainmentLevelSulfurDioxideAndSulfites string `json:"U_BOYX_Svovldioxid"`
	ContainmentLevelSesameSeeds              string `json:"U_BOYX_Sesam"`
	ContainmentLevelLupine                   string `json:"U_BOYX_Lupin"`
	ContainmentLevelMollusks                 string `json:"U_BOYX_BL"`

	// Claims
	GlutenFree  string `json:"U_BOYX_Gluten1"`
	LactoseFree string `json:"U_BOYX_Lactose"`
	Vegetarian  string `json:"U_BOYX_Vegetar"`
	Vegan       string `json:"U_BOYX_Vegan"`
	CowFree     string `json:"U_BOYX_Okse"`
	PigFree     string `json:"U_BOYX_gris"`
	GMOFree     string `json:"U_BOYX_GMO"`

	// Nutritional Information
	EnergyInkJ                    string `json:"U_BOYX_Energi"`
	EnergyInKcal                  string `json:"U_BOYX_Energik"`
	NutritionalFatValue           string `json:"U_BOYX_fedt"`
	NutritionalFattyAcidsValue    string `json:"U_BOYX_fedtsyre"`
	NutritionalCarboHydratesValue string `json:"U_BOYX_Kulhydrat"`
	NutritionalSugarValue         string `json:"U_BOYX_sukkerarter"`
	NutritionalProteinValue       string `json:"U_BOYX_Protein"`
	NutritionalSaltValue          string `json:"U_BOYX_salt"`
	ListOfIngredientsDA           string `json:"U_BOYX_varedel"`
}

type SapApiGetItemsDataReturn struct {
	Body *SapApiGetItemsDataResults
}

func SapApiGetItemsData(params SapApiQueryParams) (SapApiGetItemsDataReturn, error) {
	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticated client")
		return SapApiGetItemsDataReturn{}, err
	}

	resp, err := client.
		//DevMode().
		R().
		SetSuccessResult(SapApiGetItemsDataResults{}).
		SetQueryParams(params.AsReqParams()).
		Get("Items")
	if err != nil {
		fmt.Println(err)
		return SapApiGetItemsDataReturn{}, err
	}

	return SapApiGetItemsDataReturn{
		Body: resp.SuccessResult().(*SapApiGetItemsDataResults),
	}, nil

}

func SapApiGetItemsData_AllPages(params SapApiQueryParams) (SapApiGetItemsDataReturn, error) {
	res := SapApiGetItemsDataResults{}
	for page := 0; ; page++ {
		params.Skip = page * 20

		getItemsRes, err := SapApiGetItemsData(params)
		if err != nil {
			return SapApiGetItemsDataReturn{}, err
		}

		res.Value = append(res.Value, getItemsRes.Body.Value...)

		if getItemsRes.Body.NextLink == "" {
			break
		}
	}

	return SapApiGetItemsDataReturn{
		Body: &res,
	}, nil
}
