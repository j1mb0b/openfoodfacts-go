// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

import "encoding/json"

type Product struct {
	Id                                          string         `json:"id,omitempty"`
	Code                                        string         `json:"code,omitempty"`
	Brands                                      string         `json:"brands,omitempty"`
	BrandsTags                                  []string       `json:"brands_tags,omitempty"`
	GenericName                                 string         `json:"generic_name,omitempty"`
	GenericNameEn                               string         `json:"generic_name_en,omitempty"`
	CreatedTime                                 EpochTime      `json:"created_t,omitempty"`
	LastImageTime                               EpochTime      `json:"last_image_t,omitempty"`
	LastModifiedTime                            EpochTime      `json:"last_modified_t,omitempty"`
	CompletedTime                               EpochTime      `json:"completed_t,omitempty"`
	ImageFrontSmallURL                          URL            `json:"image_front_smallURL,omitempty"`
	ImageFrontThumbURL                          URL            `json:"image_front_thumb_url,omitempty"`
	ImageFrontURL                               URL            `json:"image_front_url,omitempty"`
	ImageIngredientsSmallURL                    URL            `json:"image_ingredients_small_url,omitempty"`
	ImageIngredientsThumbURL                    URL            `json:"image_ingredients_thumb_url,omitempty"`
	ImageIngredientsURL                         URL            `json:"image_ingredients_url,omitempty"`
	ImageNutritionSmallURL                      URL            `json:"image_nutrition_small_url,omitempty"`
	ImageNutritionThumbURL                      URL            `json:"image_nutrition_thumb_url,omitempty"`
	ImageNutritionURL                           URL            `json:"image_nutrition_url,omitempty"`
	ImageSmallURL                               URL            `json:"image_small_url,omitempty"`
	ImageThumbURL                               URL            `json:"image_thumb_url,omitempty"`
	ImageURL                                    URL            `json:"image_url,omitempty"`
	Creator                                     string         `json:"creator,omitempty"`
	Checkers                                    []interface{}  `json:"checkers,omitempty"`
	CheckersTags                                []interface{}  `json:"checkers_tags,omitempty"`
	Editors                                     []string       `json:"editors,omitempty"`
	EditorsTags                                 []string       `json:"editors_tags,omitempty"`
	Correctors                                  []string       `json:"correctors,omitempty"`
	CorrectorsTags                              []string       `json:"correctors_tags,omitempty"`
	Informers                                   []string       `json:"informers,omitempty"`
	InformersTags                               []string       `json:"informers_tags,omitempty"`
	Photographers                               []string       `json:"photographers,omitempty"`
	PhotographersTags                           []string       `json:"photographers_tags,omitempty"`
	LastEditDatesTags                           []string       `json:"last_edit_dates_tags,omitempty"`
	LastEditor                                  string         `json:"last_editor,omitempty"`
	LastImageDatesTags                          []string       `json:"last_image_dates_tags,omitempty"`
	LastModifiedBy                              string         `json:"last_modified_by,omitempty"`
	Additives                                   string         `json:"additives,omitempty"`
	AdditivesDebugTags                          []interface{}  `json:"additives_debug_tags,omitempty"`
	AdditivesNumber                             int            `json:"additives_n,omitempty"`
	AdditivesOldNumber                          int            `json:"additives_old_n,omitempty"`
	AdditivesOldTags                            []string       `json:"additives_old_tags,omitempty"`
	AdditivesPrev                               string         `json:"additives_prev,omitempty"`
	AdditivesPrevNumber                         int            `json:"additives_prev_n,omitempty"`
	AdditivesPrevTags                           []string       `json:"additives_prev_tags,omitempty"`
	AdditivesTags                               []string       `json:"additives_tags,omitempty"`
	AdditivesTagsNumber                         interface{}    `json:"additives_tags_n,omitempty"`
	Allergens                                   string         `json:"allergens,omitempty"`
	AllergensHierarchy                          []interface{}  `json:"allergens_hierarchy,omitempty"`
	AllergensTags                               []interface{}  `json:"allergens_tags,omitempty"`
	Categories                                  string         `json:"categories,omitempty"`
	CategoriesDebugTags                         []interface{}  `json:"categories_debug_tags,omitempty"`
	CategoriesHierarchy                         []string       `json:"categories_hierarchy,omitempty"`
	CategoriesPrevHierarchy                     []string       `json:"categories_prev_hierarchy,omitempty"`
	CategoriesPrevTags                          []string       `json:"categories_prev_tags,omitempty"`
	CategoriesTags                              []string       `json:"categories_tags,omitempty"`
	CitiesTags                                  []interface{}  `json:"cities_tags,omitempty"`
	CodesTags                                   []string       `json:"codes_tags,omitempty"`
	Complete                                    int            `json:"complete,omitempty"`
	Countries                                   string         `json:"countries,omitempty"`
	CountriesHierarchy                          []string       `json:"countries_hierarchy,omitempty"`
	CountriesTags                               []string       `json:"countries_tags,omitempty"`
	EmbCodes                                    string         `json:"emb_codes,omitempty"`
	EmbCodes20141016                            string         `json:"emb_codes_20141016,omitempty"`
	EmbCodesOrig                                string         `json:"emb_codes_orig,omitempty"`
	EmbCodesTags                                []interface{}  `json:"emb_codes_tags,omitempty"`
	EntryDatesTags                              []string       `json:"entry_dates_tags,omitempty"`
	EcoscoreGrade                               string         `json:"ecoscore_grade,omitempty"`
	EcoscoreTags                                []string       `json:"ecoscore_tags,omitempty"`
	ExpirationDate                              string         `json:"expiration_date,omitempty"`
	FruitsVegetablesNuts100GEstimate            json.Number    `json:"fruits-vegetables-nuts_100g_estimate,omitempty"`
	Ingredients                                 []Ingredient   `json:"ingredients,omitempty"`
	IngredientsDebug                            []interface{}  `json:"ingredients_debug,omitempty"`
	IngredientsFromOrThatMayBeFromPalmOilNumber int            `json:"ingredients_from_or_that_may_be_from_palm_oil_n,omitempty"`
	IngredientsFromPalmOilNumber                int            `json:"ingredients_from_palm_oil_n,omitempty"`
	IngredientsFromPalmOilTags                  []interface{}  `json:"ingredients_from_palm_oil_tags,omitempty"`
	IngredientsIdsDebug                         []string       `json:"ingredients_ids_debug,omitempty"`
	IngredientsN                                json.Number    `json:"ingredients_n,omitempty"`
	IngredientsNTags                            []string       `json:"ingredients_n_tags,omitempty"`
	IngredientsTags                             []string       `json:"ingredients_tags,omitempty"`
	IngredientsText                             string         `json:"ingredients_text,omitempty"`
	IngredientsTextDebug                        string         `json:"ingredients_text_debug,omitempty"`
	IngredientsTextEn                           string         `json:"ingredients_text_en,omitempty"`
	IngredientsTextWithAllergens                string         `json:"ingredients_text_with_allergens,omitempty"`
	IngredientsTextWithAllergensEn              string         `json:"ingredients_text_with_allergens_en,omitempty"`
	IngredientsThatMayBeFromPalmOilNumber       int            `json:"ingredients_that_may_be_from_palm_oil_n,omitempty"`
	IngredientsThatMayBeFromPalmOilTags         []interface{}  `json:"ingredients_that_may_be_from_palm_oil_tags,omitempty"`
	InterfaceVersionCreated                     string         `json:"interface_version_created,omitempty"`
	InterfaceVersionModified                    string         `json:"interface_version_modified,omitempty"`
	Keywords                                    []string       `json:"_keywords,omitempty"`
	Labels                                      string         `json:"labels,omitempty"`
	LabelsDebugTags                             []interface{}  `json:"labels_debug_tags,omitempty"`
	LabelsHierarchy                             []string       `json:"labels_hierarchy,omitempty"`
	LabelsPrevHierarchy                         []string       `json:"labels_prev_hierarchy,omitempty"`
	LabelsPrevTags                              []string       `json:"labels_prev_tags,omitempty"`
	LabelsTags                                  []string       `json:"labels_tags,omitempty"`
	Lang                                        string         `json:"lang,omitempty"`
	Languages                                   map[string]int `json:"languages,omitempty"`
	LanguagesCodes                              map[string]int `json:"languages_codes,omitempty"`
	LanguagesHierarchy                          []string       `json:"languages_hierarchy,omitempty"`
	LanguagesTags                               []string       `json:"languages_tags,omitempty"`
	Locale                                      string         `json:"lc,omitempty"`
	ManufacturingPlaces                         string         `json:"manufacturing_places,omitempty"`
	ManufacturingPlacesTags                     []string       `json:"manufacturing_places_tags,omitempty"`
	MaxImgId                                    string         `json:"max_imgid,omitempty"`
	NewAdditivesNumber                          int            `json:"new_additives_n,omitempty"`
	NoNutritionData                             interface{}    `json:"no_nutrition_data,omitempty"`
	NutrientLevels                              NutrientLevel  `json:"nutrient_levels,omitempty"`
	NutrientLevelsTags                          []string       `json:"nutrient_levels_tags,omitempty"`
	Nutriments                                  Nutriment      `json:"nutriments,omitempty"`
	NutritionDataPer                            string         `json:"nutrition_data_per,omitempty"`
	NutritionGradeFr                            string         `json:"nutrition_grade_fr,omitempty"`
	NutritionGrades                             string         `json:"nutrition_grades,omitempty"`
	NutritionGradesTags                         []string       `json:"nutrition_grades_tags,omitempty"`
	NutritionScoreDebug                         string         `json:"nutrition_score_debug,omitempty"`
	OriginTags                                  []string       `json:"origins_tags,omitempty"`
	Origins                                     string         `json:"origins,omitempty"`
	Packaging                                   string         `json:"packaging,omitempty"`
	PackagingTags                               []string       `json:"packaging_tags,omitempty"`
	PnnsGroups1                                 string         `json:"pnns_groups_1,omitempty"`
	PnnsGroups1Tags                             []string       `json:"pnns_groups_1_tags,omitempty"`
	PnnsGroups2                                 string         `json:"pnns_groups_2,omitempty"`
	PnnsGroups2Tags                             []string       `json:"pnns_groups_2_tags,omitempty"`
	ProductName                                 string         `json:"product_name,omitempty"`
	ProductNameEn                               string         `json:"product_name_en,omitempty"`
	PurchasePlaces                              string         `json:"purchase_places,omitempty"`
	PurchasePlacesTags                          []interface{}  `json:"purchase_places_tags,omitempty"`
	Quantity                                    string         `json:"quantity,omitempty"`
	Rev                                         json.Number    `json:"rev,omitempty"`
	ScansNumber                                 int            `json:"scans_n,omitempty"`
	ServingQuantity                             json.Number    `json:"serving_quantity,omitempty"`
	ServingSize                                 string         `json:"serving_size,omitempty"`
	SortKey                                     int            `json:"sortkey,omitempty"`
	States                                      string         `json:"states,omitempty"`
	StatesHierarchy                             []string       `json:"states_hierarchy,omitempty"`
	StatesTags                                  []string       `json:"states_tags,omitempty"`
	Stores                                      string         `json:"stores,omitempty"`
	StoresTags                                  []interface{}  `json:"stores_tags,omitempty"`
	Traces                                      string         `json:"traces,omitempty"`
	TracesHierarchy                             []string       `json:"traces_hierarchy,omitempty"`
	TracesTags                                  []string       `json:"traces_tags,omitempty"`
	UniqueScansNumber                           int            `json:"unique_scans_n,omitempty"`
	UnknownNutrientsTags                        []interface{}  `json:"unknown_nutrients_tags,omitempty"`
	UpdateKey                                   string         `json:"update_key,omitempty"`
}
