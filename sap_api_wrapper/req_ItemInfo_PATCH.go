package sap_api_wrapper

import (
	"fmt"
	"strings"
	"time"
)

type SapApiItemsDataBody struct {
	ItemCode      string    `json:"ItemCode"`
	UpdateNUT     string    `json:"U_CCF_Update_NUT"`
	NutUpdateDate time.Time `json:"U_CCF_NUT_UpdateDate"`

	// Lists of ingredients
	IngredientsScandinavian string `json:"U_CCF_Ingrediens_DA_SE_NO"`
	IngredientsDanish       string `json:"U_CCF_Ingrediens_DA"`
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
	ContainmentLevelSesameSeeds              string `json:"U_BOYX_Sesam"`
	ContainmentLevelSulfurDioxideAndSulfites string `json:"U_BOYX_Svovldioxid"`
	ContainmentLevelLupine                   string `json:"U_BOYX_Lupin"`
	ContainmentLevelMollusks                 string `json:"U_BOYX_BL"`

	// Claims
	GlutenFree  string `json:"U_BOYX_Gluten1"`
	LactoseFree string `json:"U_BOYX_Lactose"`
	Vegetarian  string `json:"U_BOXY_Vegetar"`
	Vegan       string `json:"U_BOXY_Vegan"`
	CowFree     string `json:"U_BOYX_Oske"`
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
}

type SapApiErrorResult struct {
	ErrorMessage struct {
		Code    int `json:"code"`
		Message struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"message"`
	} `json:"error"`
}

// Takes a SAP Item and patches it in SAP
func SetItemData(item SapApiItemsData) error {

	var itemDataBody SapApiItemsDataBody
	itemDataBody.ItemCode = item.ItemCode
	itemDataBody.UpdateNUT = "N"
	itemDataBody.NutUpdateDate = time.Now().UTC()
	itemDataBody.IngredientsScandinavian = item.IngredientsScandinavian
	itemDataBody.IngredientsDanish = item.IngredientsDanish
	itemDataBody.IngredientsFinnish = item.IngredientsFinnish
	itemDataBody.IngredientsEnglish = item.IngredientsEnglish
	itemDataBody.IngredientsGerman = item.IngredientsGerman
	itemDataBody.IngredientsDutch = item.IngredientsDutch
	itemDataBody.IngredientsFrench = item.IngredientsFrench
	itemDataBody.IngredientsPortuguese = item.IngredientsPortuguese
	itemDataBody.IngredientsItalian = item.IngredientsItalian
	itemDataBody.IngredientsSpanish = item.IngredientsSpanish
	itemDataBody.EnergyInkJ = strings.Replace(fmt.Sprint(item.EnergyInkJ), ".", ",", -1)
	itemDataBody.EnergyInKcal = strings.Replace(fmt.Sprint(item.EnergyInKcal), ".", ",", -1)
	itemDataBody.NutritionalFatValue = strings.Replace(fmt.Sprint(item.NutritionalFatValue), ".", ",", -1)
	itemDataBody.NutritionalFattyAcidsValue = strings.Replace(fmt.Sprint(item.NutritionalFattyAcidsValue), ".", ",", -1)
	itemDataBody.NutritionalCarboHydratesValue = strings.Replace(fmt.Sprint(item.NutritionalCarboHydratesValue), ".", ",", -1)
	itemDataBody.NutritionalSugarValue = strings.Replace(fmt.Sprint(item.NutritionalSugarValue), ".", ",", -1)
	itemDataBody.NutritionalProteinValue = strings.Replace(fmt.Sprint(item.NutritionalProteinValue), ".", ",", -1)
	itemDataBody.NutritionalSaltValue = strings.Replace(fmt.Sprint(item.NutritionalSaltValue), ".", ",", -1)

	itemDataBody.ContainmentLevelGluten = item.ContainmentLevelGluten
	itemDataBody.ContainmentLevelCrustacea = item.ContainmentLevelCrustacea
	itemDataBody.ContainmentLevelEgg = item.ContainmentLevelEgg
	itemDataBody.ContainmentLevelFish = item.ContainmentLevelFish
	itemDataBody.ContainmentLevelPeanut = item.ContainmentLevelPeanut
	itemDataBody.ContainmentLevelSoy = item.ContainmentLevelSoy
	itemDataBody.ContainmentLevelMilk = item.ContainmentLevelMilk
	itemDataBody.ContainmentLevelAlmonds = item.ContainmentLevelAlmonds
	itemDataBody.ContainmentLevelHazelnut = item.ContainmentLevelHazelnut
	itemDataBody.ContainmentLevelWalnut = item.ContainmentLevelWalnut
	itemDataBody.ContainmentLevelCashew = item.ContainmentLevelCashew
	itemDataBody.ContainmentLevelPecan = item.ContainmentLevelPecan
	itemDataBody.ContainmentLevelBrazilNut = item.ContainmentLevelBrazilNut
	itemDataBody.ContainmentLevelPistachio = item.ContainmentLevelPistachio
	itemDataBody.ContainmentLevelQueenslandNut = item.ContainmentLevelQueenslandNut
	itemDataBody.ContainmentLevelCelery = item.ContainmentLevelCelery

	itemDataBody.ContainmentLevelMustard = item.ContainmentLevelMustard
	itemDataBody.ContainmentLevelSulfurDioxideAndSulfites = item.ContainmentLevelSulfurDioxideAndSulfites
	itemDataBody.ContainmentLevelSesameSeeds = item.ContainmentLevelSesameSeeds
	itemDataBody.ContainmentLevelLupine = item.ContainmentLevelLupine
	itemDataBody.ContainmentLevelMollusks = item.ContainmentLevelMollusks

	if itemDataBody.ContainmentLevelGluten == "Free from" {
		itemDataBody.GlutenFree = "Y"
	}
	if itemDataBody.ContainmentLevelMilk == "Free from" {
		itemDataBody.LactoseFree = "Y"
		itemDataBody.Vegan = "Y"
	}

	itemDataBody.Vegetarian = "Y"
	itemDataBody.CowFree = "Y"
	itemDataBody.PigFree = "Y"
	itemDataBody.GMOFree = "Y"

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
		SetErrorResult(SapApiErrorResult{}).
		SetHeader("Content-Type", "application/json").
		SetBody(itemDataBody).
		Patch(fmt.Sprintf("Items('%v')", item.ItemCode))
	if err != nil {
		return err
	}
	if resp.IsErrorState() {
		errorResponse := resp.ErrorResult().(*SapApiErrorResult)
		if resp.StatusCode == 400 {
			fmt.Printf("resp is errorcode 400, response:%v\n", errorResponse.ErrorMessage.Message.Value)
		} else {
			fmt.Printf("resp is err statusCode: %v. Dump: %v\n", resp.StatusCode, resp.Dump())
			return nil
		}
	}

	return nil
}
