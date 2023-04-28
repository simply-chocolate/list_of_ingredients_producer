package utils

// Gets the bill of materials for an itemcode and adds all the raw materials to the list of raw materials
// Also returns a list of HF's and FV's that we need to check the bill of materials for if there are any
func GetBillOfMaterials(itemCode string, billOfMaterials BillOfMaterials, QuantityOnFather float64) (BillOfMaterials, []ItemCodeAndQuantity, error) {
	HFs := []ItemCodeAndQuantity{}

	// While has HF is false, we will keep calling the function and going deeper into the bill of materials
	currentBillOfMaterials, err := GetBillOfMaterialFromSap(itemCode)
	if err != nil {
		panic(err)
	}
	for _, currentBillOfMaterial := range currentBillOfMaterials.Value {
		if currentBillOfMaterial.Type == "Råvare" {
			// multiply the quantity with the quantity of the father
			currentBillOfMaterial.Quantity = RoundWithPrecision(currentBillOfMaterial.Quantity*QuantityOnFather, 6)
			billOfMaterials = append(billOfMaterials, currentBillOfMaterial)

		} else if currentBillOfMaterial.Type == "HF" {
			var HF ItemCodeAndQuantity
			HF.ItemCode = currentBillOfMaterial.ItemCode
			HF.Quantity = RoundWithPrecision(currentBillOfMaterial.Quantity*QuantityOnFather, 6)

			HFs = append(HFs, HF)
		} else if currentBillOfMaterial.Type == "Færdigvare" {
			var FV ItemCodeAndQuantity
			FV.ItemCode = currentBillOfMaterial.ItemCode
			FV.Quantity = RoundWithPrecision(currentBillOfMaterial.Quantity*QuantityOnFather, 6)

			HFs = append(HFs, FV)
		}
	}

	return billOfMaterials, HFs, nil
}
