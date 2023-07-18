package utils

import (
	"fmt"
	"list_of_ingredients_producer/sap_api_wrapper"
	"strconv"
	"strings"
)

// Takes the map of raw materials and calculates the nutritional values for the sales item
// Returns a map of nutritional values
func CalculateNutritionalValues(
	allRawMaterialsMap map[string]map[string]string,
	ingredientsOnProduct []RawMaterial,
	salesItem sap_api_wrapper.SapApiItemsData,
	totalQuantityForItem float64) (map[string]string, error) {
	var err error
	nutritionalValuesMap := map[string]float64{
		"EnergyKJ":      0.0,
		"EnergyKcal":    0.0,
		"Fat":           0.0,
		"FattyAcids":    0.0,
		"CarboHydrates": 0.0,
		"Sugar":         0.0,
		"Protein":       0.0,
		"Salt":          0.0,
	}

	for _, ingredient := range ingredientsOnProduct {
		ingredientFromMap, exists := allRawMaterialsMap[ingredient.ItemCode]
		if !exists {
			return map[string]string{}, fmt.Errorf("error getting raw material from map with itemcode: %v", ingredient.ItemCode)
		}

		nutrientFactor := ingredient.Quantity / totalQuantityForItem
		nutritionalValuesMap, err = handleAppendingNutritionalValues(ingredient.ItemCode, ingredientFromMap, nutritionalValuesMap, nutrientFactor)
		if err != nil {
			return map[string]string{}, err
		}
	}

	nutritionalStringMap := make(map[string]string)
	nutritionalStringMap["EnergyKJ"] = FindCorrectPrecision(nutritionalValuesMap["EnergyKJ"])
	nutritionalStringMap["EnergyKcal"] = FindCorrectPrecision(nutritionalValuesMap["EnergyKcal"])
	nutritionalStringMap["Fat"] = FindCorrectPrecision(nutritionalValuesMap["Fat"])
	nutritionalStringMap["FattyAcids"] = FindCorrectPrecision(nutritionalValuesMap["FattyAcids"])
	nutritionalStringMap["CarboHydrates"] = FindCorrectPrecision(nutritionalValuesMap["CarboHydrates"])
	nutritionalStringMap["Sugar"] = FindCorrectPrecision(nutritionalValuesMap["Sugar"])
	nutritionalStringMap["Protein"] = FindCorrectPrecision(nutritionalValuesMap["Protein"])
	nutritionalStringMap["Salt"] = FindCorrectPrecision(nutritionalValuesMap["Salt"])

	return nutritionalStringMap, nil
}

func handleAppendingNutritionalValues(
	itemCodeRawMaterial string,
	rawMaterial map[string]string,
	nutritionalValuesMap map[string]float64,
	nutriantFactor float64,
) (map[string]float64, error) {

	energiKJ, err := strconv.ParseFloat(strings.Replace(rawMaterial["EnergyKJ"], ",", ".", -1), 64)
	if err != nil {
		return map[string]float64{}, fmt.Errorf("error parsing EnergyKJ to float with value: %v for rawMaterials: %v", err, itemCodeRawMaterial)
	}
	energiKcal, err := strconv.ParseFloat(strings.Replace(rawMaterial["EnergyKcal"], ",", ".", -1), 64)
	if err != nil {
		return map[string]float64{}, fmt.Errorf("error parsing EnergyKcal to float with value: %v for rawMaterials: %v", err, itemCodeRawMaterial)
	}
	fat, err := strconv.ParseFloat(strings.Replace(rawMaterial["Fat"], ",", ".", -1), 64)
	if err != nil {
		return map[string]float64{}, fmt.Errorf("error parsing Fat to float with value: %v for rawMaterials: %v", err, itemCodeRawMaterial)
	}
	fattyAcids, err := strconv.ParseFloat(strings.Replace(rawMaterial["FattyAcids"], ",", ".", -1), 64)
	if err != nil {
		return map[string]float64{}, fmt.Errorf("error parsing FattyAcids to float with value: %v for rawMaterials: %v", err, itemCodeRawMaterial)
	}
	carboHydrates, err := strconv.ParseFloat(strings.Replace(rawMaterial["CarboHydrates"], ",", ".", -1), 64)
	if err != nil {
		return map[string]float64{}, fmt.Errorf("error parsing CarboHydrates to float with value: %v for rawMaterials: %v", err, itemCodeRawMaterial)
	}
	sugar, err := strconv.ParseFloat(strings.Replace(rawMaterial["Sugar"], ",", ".", -1), 64)
	if err != nil {
		return map[string]float64{}, fmt.Errorf("error parsing Sugar to float with value: %v for rawMaterials: %v", err, itemCodeRawMaterial)
	}
	protein, err := strconv.ParseFloat(strings.Replace(rawMaterial["Protein"], ",", ".", -1), 64)
	if err != nil {
		return map[string]float64{}, fmt.Errorf("error parsing Protein to float with value: %v for rawMaterials: %v", err, itemCodeRawMaterial)
	}
	/*
		salt, err := strconv.ParseFloat(strings.Replace(rawMaterial["Salt"], ",", ".", -1), 64)
		if err != nil {
			return map[string]float64{}, fmt.Errorf("error parsing Salt to float with value: %v for rawMaterials: %v", err, itemCodeRawMaterial)
		}
	*/
	salt, err := strconv.ParseFloat(strings.Replace(rawMaterial["Salt"], ",", ".", -1), 64)
	if err != nil {
		return map[string]float64{}, fmt.Errorf("error parsing Salt to float with value: %v for rawMaterials: %v", err, itemCodeRawMaterial)
	}

	nutritionalValuesMap["EnergyKJ"] += energiKJ * nutriantFactor
	nutritionalValuesMap["EnergyKcal"] += energiKcal * nutriantFactor
	nutritionalValuesMap["Fat"] += fat * nutriantFactor
	nutritionalValuesMap["FattyAcids"] += fattyAcids * nutriantFactor
	nutritionalValuesMap["CarboHydrates"] += carboHydrates * nutriantFactor
	nutritionalValuesMap["Sugar"] += sugar * nutriantFactor
	nutritionalValuesMap["Protein"] += protein * nutriantFactor
	nutritionalValuesMap["Salt"] += salt * nutriantFactor

	return nutritionalValuesMap, nil
}
