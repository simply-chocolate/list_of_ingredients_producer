﻿# list_of_ingredients_producer

Created by Jedikrigeren for simply-chocolate

The purpose of this script is to look for all the items requesting an updated list of ingredients in SAP.
It looks for the SAP field: `U_CCF_Update_NUT` that has the value `Y`.

Common user errors that does not return an error:
  - The raw materials on the child items does not have the type "Råvare". -> This leads to the raw material not being properly used in the creation of the list of ingredients.
  - The items does not contain the item numbers for the claimed items in the correct field. -> This leads to the list of ingredients not having the correct percentage marks
  - The raw materials are missing other vital information (e.g. list of ingredients, nutritional information or contamination information). -> This can lead to error in contamination information




