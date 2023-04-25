package utils

// In here we need to find the raw materials for the recipe

// We need to take an argument that is the itemcode
// Then we will use the itemcode to find the bill of materials
// Then we will check the bill of materials to see if it is a raw material
// If the typing is raw material, we will add it to the list of raw materials for the itemcode
// 			We are going to store the list of ingredients for the raw material in a map containing each language
// If the typing is not raw material, we will call the function again with the new itemcode
// We will need to return a list of raw materials for the itemcode
