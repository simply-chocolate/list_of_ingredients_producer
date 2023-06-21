package utils

import (
	"list_of_ingredients_producer/sap_api_wrapper"
)

// Takes a containmentMap and the Raw Material containments and adjusts the containment map if the Raw Material has any allergens with greater contamination
// returns the adjusted containment map
func FindAllAllergenContaminations(containMentMap map[string]string, RawMaterialInfo map[string]string) map[string]string {
	containMentMap["Gluten"] = checkIfContainmentIsGreater(containMentMap["Gluten"], RawMaterialInfo["Gluten"])
	containMentMap["Crustacea"] = checkIfContainmentIsGreater(containMentMap["Crustacea"], RawMaterialInfo["Crustacea"])
	containMentMap["Egg"] = checkIfContainmentIsGreater(containMentMap["Egg"], RawMaterialInfo["Egg"])
	containMentMap["Fish"] = checkIfContainmentIsGreater(containMentMap["Fish"], RawMaterialInfo["Fish"])
	containMentMap["Peanut"] = checkIfContainmentIsGreater(containMentMap["Peanut"], RawMaterialInfo["Peanut"])
	containMentMap["Soy"] = checkIfContainmentIsGreater(containMentMap["Soy"], RawMaterialInfo["Soy"])
	containMentMap["Milk"] = checkIfContainmentIsGreater(containMentMap["Milk"], RawMaterialInfo["Milk"])
	containMentMap["Almonds"] = checkIfContainmentIsGreater(containMentMap["Almonds"], RawMaterialInfo["Almonds"])
	containMentMap["Hazelnut"] = checkIfContainmentIsGreater(containMentMap["Hazelnut"], RawMaterialInfo["Hazelnut"])
	containMentMap["Walnut"] = checkIfContainmentIsGreater(containMentMap["Walnut"], RawMaterialInfo["Walnut"])
	containMentMap["Cashew"] = checkIfContainmentIsGreater(containMentMap["Cashew"], RawMaterialInfo["Cashew"])
	containMentMap["Pecan"] = checkIfContainmentIsGreater(containMentMap["Pecan"], RawMaterialInfo["Pecan"])
	containMentMap["BrazilNut"] = checkIfContainmentIsGreater(containMentMap["BrazilNut"], RawMaterialInfo["BrazilNut"])
	containMentMap["Pistachio"] = checkIfContainmentIsGreater(containMentMap["Pistachio"], RawMaterialInfo["Pistachio"])
	containMentMap["QueenslandNut"] = checkIfContainmentIsGreater(containMentMap["QueenslandNut"], RawMaterialInfo["QueenslandNut"])
	containMentMap["Celery"] = checkIfContainmentIsGreater(containMentMap["Celery"], RawMaterialInfo["Celery"])

	return containMentMap
}

// Takes a containmentMap and the Sales Item containments and adjusts the containment map if the Sales Item has any allergens with greater contamination
// returns the adjusted containment map
func FindAllAllergenContaminationsSalesItem(containMentMap map[string]string, salesItem sap_api_wrapper.SapApiItemsData) map[string]string {
	containMentMap["Gluten"] = checkIfContainmentIsGreater(containMentMap["Gluten"], salesItem.ContainmentLevelGluten)
	containMentMap["Crustacea"] = checkIfContainmentIsGreater(containMentMap["Crustacea"], salesItem.ContainmentLevelCrustacea)
	containMentMap["Egg"] = checkIfContainmentIsGreater(containMentMap["Egg"], salesItem.ContainmentLevelEgg)
	containMentMap["Fish"] = checkIfContainmentIsGreater(containMentMap["Fish"], salesItem.ContainmentLevelFish)
	containMentMap["Peanut"] = checkIfContainmentIsGreater(containMentMap["Peanut"], salesItem.ContainmentLevelPeanut)
	containMentMap["Soy"] = checkIfContainmentIsGreater(containMentMap["Soy"], salesItem.ContainmentLevelSoy)
	containMentMap["Milk"] = checkIfContainmentIsGreater(containMentMap["Milk"], salesItem.ContainmentLevelMilk)
	containMentMap["Almonds"] = checkIfContainmentIsGreater(containMentMap["Almonds"], salesItem.ContainmentLevelAlmonds)
	containMentMap["Hazelnut"] = checkIfContainmentIsGreater(containMentMap["Hazelnut"], salesItem.ContainmentLevelHazelnut)
	containMentMap["Walnut"] = checkIfContainmentIsGreater(containMentMap["Walnut"], salesItem.ContainmentLevelWalnut)
	containMentMap["Cashew"] = checkIfContainmentIsGreater(containMentMap["Cashew"], salesItem.ContainmentLevelCashew)
	containMentMap["Pecan"] = checkIfContainmentIsGreater(containMentMap["Pecan"], salesItem.ContainmentLevelPecan)
	containMentMap["BrazilNut"] = checkIfContainmentIsGreater(containMentMap["BrazilNut"], salesItem.ContainmentLevelBrazilNut)
	containMentMap["Pistachio"] = checkIfContainmentIsGreater(containMentMap["Pistachio"], salesItem.ContainmentLevelPistachio)
	containMentMap["QueenslandNut"] = checkIfContainmentIsGreater(containMentMap["QueenslandNut"], salesItem.ContainmentLevelQueenslandNut)
	containMentMap["Celery"] = checkIfContainmentIsGreater(containMentMap["Celery"], salesItem.ContainmentLevelCelery)

	return containMentMap
}

// Checks which contamination level is greater between to levels of contamination
// Returns the level of contamination that is greater
func checkIfContainmentIsGreater(currentContainment string, rawMaterialContainment string) string {
	containMentChecker := map[string]int{
		"Free from":             1,
		"May contain traces of": 2,
		"In product":            3,
	}

	if containMentChecker[currentContainment] < containMentChecker[rawMaterialContainment] {
		return rawMaterialContainment
	} else {
		return currentContainment
	}
}

// Takes a containment map and a language code
// returns a string of the allergens that the product may contain traces of in the given language
func createStringOfTraceContamination(containMentMap map[string]string, languageCode string) string {
	containmentString := getStartOfMayContainMap()[languageCode]
	tracesSlice := []string{}
	nutMap := getNutMap()
	hasNuts := false

	// First we need to make a slice of all the "May contain traces of" allergens that are not nuts (except peanuts)
	for allergen, containment := range containMentMap {

		if _, exists := nutMap[allergen]; exists {
			hasNuts = true
			continue
		}

		// If containment is "May contain traces of" and it is not a nut, we need to add it to the slice
		if containment == "May contain traces of" {
			tracesSlice = append(tracesSlice, allergen)
		}
	}

	// If theres trace of any nuts in the product, we add it to the slice
	if hasNuts {
		tracesSlice = append(tracesSlice, "Nuts")
	}

	// Then we need to add the allergens to the string
	allergenMap := getAllergenMap()

	for i, allergen := range tracesSlice {

		if i == len(tracesSlice)-2 {
			containmentString += allergenMap[allergen][languageCode] + getAndMap()[languageCode]
		} else if i == len(tracesSlice)-1 {
			containmentString += allergenMap[allergen][languageCode]
		} else {
			containmentString += allergenMap[allergen][languageCode] + ", "
		}
	}

	return containmentString + getEndOfMayContainMap()[languageCode]
}

// Return a map of the allergens in all languages with language code as key
func getAllergenMap() map[string]map[string]string {
	allergenMap := map[string]map[string]string{
		"Gluten":    {"DA_SE_NO": "gluten", "FI": "gluteeni", "EN": "gluten", "DE": "Gluten", "NL": "gluten", "FR": "gluten", "PT": "glúten", "IT": "glutine", "ES": "gluten"},
		"Crustacea": {"DA_SE_NO": "krebsdyr/skaldjur/krepsdyr", "FI": "äyriäiset", "EN": "crustaceans", "DE": "Krebstiere", "NL": "schaaldieren", "FR": "crustacés", "PT": "crustáceos", "IT": "crostacei", "ES": "crustáceos"},
		"Egg":       {"DA_SE_NO": "æg", "FI": "muna", "EN": "egg", "DE": "Ei", "NL": "ei", "FR": "oeuf", "PT": "ovo", "IT": "uovo", "ES": "huevo"},
		"Fish":      {"DA_SE_NO": "fisk", "FI": "kala", "EN": "fish", "DE": "Fisch", "NL": "vis", "FR": "poisson", "PT": "peixe", "IT": "pesce", "ES": "pescado"},
		"Peanut":    {"DA_SE_NO": "peanuts/jordnötter/peanøtt", "FI": "maapähkinä", "EN": "peanuts", "DE": "Erdnüsse", "NL": "pinda", "FR": "arachide", "PT": "amendoim", "IT": "arachidi", "ES": "cacahuete"},
		"Soy":       {"DA_SE_NO": "soja/soya", "FI": "soija", "EN": "soy", "DE": "soja", "NL": "soja", "FR": "soja", "PT": "soja", "IT": "soia", "ES": "soja"},
		"Milk":      {"DA_SE_NO": "mælk/mjölk/melk", "FI": "maito", "EN": "milk", "DE": "Milch", "NL": "melk", "FR": "lait", "PT": "leite", "IT": "latte", "ES": "leche"},
		// TODO: These nuts are just gonna be called "nuts" in the allergen list.
		// TODO: If we want this changed, we need to translate the DA_SE_NO better, its just in danish right now
		//"Almonds":        {"DA_SE_NO": "Mandler", "FI": "Manteli", "EN": "Almonds", "DE": "Mandeln", "NL": "Amandel", "FR": "Amande", "PT": "Amêndoa", "IT": "Mandorle", "ES": "Almendra"},
		//"Hazelnut":       {"DA_SE_NO": "Hasselnødder", "FI": "Pähkinä", "EN": "Hazelnuts", "DE": "Haselnüsse", "NL": "Hazelnoten", "FR": "Noisette", "PT": "Avelã", "IT": "Nocciole", "ES": "Avellana"},
		//"Walnut":         {"DA_SE_NO": "Valnødder", "FI": "Saksanpähkinä", "EN": "Walnuts", "DE": "Walnüsse", "NL": "Walnoten", "FR": "Noix", "PT": "Noz", "IT": "Noci", "ES": "Nuez"},
		//"Cashew":         {"DA_SE_NO": "Cashewnødder", "FI": "Cashewpähkinä", "EN": "Cashews", "DE": "Cashewnüsse", "NL": "Cashewnoten", "FR": "Noix de cajou", "PT": "Caju", "IT": "Anacardi", "ES": "Anacardo"},
		//"Pecan":          {"DA_SE_NO": "Pekannødder", "FI": "Pekaanipähkinä", "EN": "Pecan nuts", "DE": "Pekannüsse", "NL": "Pecannoten", "FR": "Noix de pécan", "PT": "Noz-pecã", "IT": "Noci pecan", "ES": "Nuez pecana"},
		//"BrazilNut":      {"DA_SE_NO": "Paranødder", "FI": "Parapähkinä", "EN": "Brazil nuts", "DE": "Paranüsse", "NL": "Paranoten", "FR": "Noix du Brésil", "PT": "Noz do Brasil", "IT": "Noci del Brasile", "ES": "Nuez de Brasil"},
		//"Pistachio":      {"DA_SE_NO": "Pistacienødder", "FI": "Pistaasipähkinä", "EN": "Pistachios", "DE": "Pistazien", "NL": "Pistachenoten", "FR": "Pistache", "PT": "Pistácio", "IT": "Pistacchi", "ES": "Pistacho"},
		//"QueenslandNut":  {"DA_SE_NO": "Queenslandnødder", "FI": "Queenslandpähkinä", "EN": "Queensland nuts", "DE": "Queenslandnüsse", "NL": "Queenslandnoten", "FR": "Noix du Queensland", "PT": "Noz de Queensland", "IT": "Noci del Queensland", "ES": "Nuez de Queensland"},
		"Celery":         {"DA_SE_NO": "selleri/", "FI": "selleri", "EN": "celery", "DE": "Sellerie", "NL": "selderij", "FR": "céleri", "PT": "aipo", "IT": "sedano", "ES": "apio"},
		"Mustard":        {"DA_SE_NO": "sennep/senap", "FI": "sinappi", "EN": "mustard", "DE": "Senf", "NL": "mosterd", "FR": "moutarde", "PT": "mostarda", "IT": "senape", "ES": "mostaza"},
		"Sesame":         {"DA_SE_NO": "sesam", "FI": "seesami", "EN": "sesame", "DE": "Sesam", "NL": "sesam", "FR": "sésame", "PT": "sésamo", "IT": "sesamo", "ES": "sésamo"},
		"SulphurDioxide": {"DA_SE_NO": "svovldioxid/svaveldioxid/svoveldioksid/", "FI": "rikkidioksidi", "EN": "sulphur dioxide", "DE": "Schwefeldioxid", "NL": "zwaveldioxide", "FR": "dioxyde de soufre", "PT": "dióxido de enxofre", "IT": "anidride solforosa", "ES": "dióxido de azufre"},
		"Lupin":          {"DA_SE_NO": "lupin", "FI": "lupiini", "EN": "lupin", "DE": "Lupine", "NL": "lupine", "FR": "lupin", "PT": "lupino", "IT": "lupino", "ES": "lupino"},
		"Molluscs":       {"DA_SE_NO": "bløddyr/mollusker/", "FI": "nilviäiset", "EN": "molluscs", "DE": "Weichtiere", "NL": "weekdieren", "FR": "mollusques", "PT": "moluscos", "IT": "molluschi", "ES": "moluscos"},
		"Nuts":           {"DA_SE_NO": "nødder/nötter/nøtter", "FI": "pähkinät", "EN": "nuts", "DE": "Nüsse", "NL": "noten", "FR": "noix", "PT": "nozes", "IT": "noci", "ES": "nueces"},
	}

	return allergenMap
}

// Returns a map of which nuts are okay to write in the list of ingredients as "Nuts" instead of their actual name, if they are just traces
func getNutMap() map[string]string {
	nutmap := map[string]string{
		"Almonds":       "",
		"Hazelnut":      "",
		"Walnut":        "",
		"Cashew":        "",
		"Pecan":         "",
		"BrazilNut":     "",
		"Pistachio":     "",
		"QueenslandNut": "",
	}
	return nutmap
}

// Returns a map of the start of the "May contain traces of" in all languages with language code as key
func getStartOfMayContainMap() map[string]string {
	mayContainTracesOfMap := map[string]string{
		"DA_SE_NO": "Kan indeholde/innehålla spor/spår af/av ",
		"FI":       "Saattaa sisältää jäämiä ",
		"EN":       "May contain traces of ",
		"DE":       "Kann Spuren von ",
		"NL":       "Kan sporen bevatten van ",
		"FR":       "Peut contenir des traces ",
		"PT":       "Pode conter vestígios ",
		"IT":       "Può contenere tracce ",
		"ES":       "Puede contener trazas ",
	}
	return mayContainTracesOfMap
}

// Returns a map of the end of the "May contain traces of" in all languages with language code as key
// Only the german one has something at the end
func getEndOfMayContainMap() map[string]string {
	mayContainTracesOfMap := map[string]string{
		"DA_SE_NO": ".",
		"FI":       ".",
		"EN":       ".",
		"DE":       " enhalten.",
		"NL":       ".",
		"FR":       ".",
		"PT":       ".",
		"IT":       ".",
		"ES":       ".",
	}
	return mayContainTracesOfMap
}

// Returns a map of the differrent languages way of saying "and"
func getAndMap() map[string]string {
	andMap := map[string]string{
		"DA_SE_NO": " og ",
		"FI":       " ja ",
		"EN":       " and ",
		"DE":       " und ",
		"NL":       " en ",
		"FR":       " et ",
		"PT":       " e ",
		"IT":       " e ",
		"ES":       " y ",
	}
	return andMap
}
