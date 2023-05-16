package sap_api_wrapper

import (
	"fmt"
)

type SapApiItemsDataBody struct {
	ItemCode string `json:"ItemCode"`

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
	/*
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
	*/

	// Nutritional Information
	/*
		EnergyInkJ                    string `json:"U_BOYX_Energi"`
		EnergyInKcal                  string `json:"U_BOYX_Energik"`
		NutritionalFatValue           string `json:"U_BOYX_fedt"`
		NutritionalFattyAcidsValue    string `json:"U_BOYX_fedtsyre"`
		NutritionalCarboHydratesValue string `json:"U_BOYX_Kulhydrat"`
		NutritionalSugarValue         string `json:"U_BOYX_sukkerarter"`
		NutritionalProteinValue       string `json:"U_BOYX_Protein"`
		NutritionalSaltValue          string `json:"U_BOYX_salt"`
		ListOfIngredientsDA           string `json:"U_BOYX_varedel"`
	*/
}

// Takes a SAP Item and patches it in SAP
func SetItemData(item SapApiItemsData) error {
	var itemDataBody SapApiItemsDataBody
	itemDataBody.ItemCode = item.ItemCode
	itemDataBody.IngredientsScandinavian = item.IngredientsScandinavian
	itemDataBody.IngredientsFinnish = item.IngredientsFinnish
	itemDataBody.IngredientsEnglish = item.IngredientsEnglish
	itemDataBody.IngredientsGerman = item.IngredientsGerman
	itemDataBody.IngredientsDutch = item.IngredientsDutch
	itemDataBody.IngredientsFrench = item.IngredientsFrench
	itemDataBody.IngredientsPortuguese = item.IngredientsPortuguese
	itemDataBody.IngredientsItalian = item.IngredientsItalian
	itemDataBody.IngredientsSpanish = item.IngredientsSpanish

	client, err := GetSapApiAuthClient()
	if err != nil {
		fmt.Println("Error getting an authenticaed client")
		return err
	}

	resp, err := client.
		//DevMode().
		R().
		EnableDump().
		SetSuccessResult(SapApiPostLoginResult{}).
		SetHeader("Content-Type", "application/json").
		SetBody(itemDataBody).
		Patch(fmt.Sprintf("Items('%v')", item.ItemCode))
	if err != nil {
		return err
	}
	if resp.IsErrorState() {
		fmt.Printf("resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
		return resp.Err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("unexpected errorcode when patching the items endpoint. Itemcode:%v. StatusCode:%v", itemDataBody.ItemCode, resp.StatusCode)
	}

	return nil
}
