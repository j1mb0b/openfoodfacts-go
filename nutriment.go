// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

// Nutriment is the collection of all retruned nutriment associated with the product
type Nutriment struct {
	Salt                    float64 `json:"salt,omitempty"`
	Salt100G                float64 `json:"salt_100g,omitempty"`
	SaltValue               float64 `json:"salt_value,omitempty"`
	SaltServing             float64 `json:"salt_serving,omitempty"`
	SaltUnit                string  `json:"salt_unit,omitempty"`
	Sugars100G              float64 `json:"sugars_100g,omitempty"`
	Sugars                  float64 `json:"sugars,omitempty"`
	SugarsUnit              string  `json:"sugars_unit,omitempty"`
	SugarsServing           float64 `json:"sugars_serving,omitempty"`
	SugarsValue             float64 `json:"sugars_value,omitempty"`
	Iron                    float64 `json:"iron,omitempty"`
	IronValue               float64 `json:"iron_value,omitempty"`
	IronLabel               string  `json:"iron_label,omitempty"`
	IronUnit                string  `json:"iron_unit,omitempty"`
	Iron100G                float64 `json:"iron_100g,omitempty"`
	IronServing             float64 `json:"iron_serving,omitempty"`
	CalciumUnit             string  `json:"calcium_unit,omitempty"`
	CalciumServing          float64 `json:"calcium_serving,omitempty"`
	Calcium                 float64 `json:"calcium,omitempty"`
	CalciumValue            float64 `json:"calcium_value,omitempty"`
	CalciumLabel            string  `json:"calcium_label,omitempty"`
	Calcium100G             float64 `json:"calcium_100g,omitempty"`
	Cholesterol100G         float64 `json:"cholesterol_100g,omitempty"`
	SaturatedFat100G        float64 `json:"saturated-fat_100g,omitempty"`
	SaturatedFatServing     float64 `json:"saturated-fat_serving,omitempty"`
	SaturatedFat            float64 `json:"saturated-fat,omitempty"`
	SaturatedFatValue       float64 `json:"saturated-fat_value,omitempty"`
	SaturatedFatUnit        string  `json:"saturated-fat_unit,omitempty"`
	Fat100G                 float64 `json:"fat_100g,omitempty"`
	FatServing              float64 `json:"fat_serving,omitempty"`
	FatValue                float64 `json:"fat_value,omitempty"`
	FatUnit                 string  `json:"fat_unit,omitempty"`
	Fat                     float64 `json:"fat,omitempty"`
	TransFatLabel           string  `json:"trans-fat_label,omitempty"`
	TransFatUnit            string  `json:"trans-fat_unit,omitempty"`
	TransFat                float64 `json:"trans-fat,omitempty"`
	TransFat100G            float64 `json:"trans-fat_100g,omitempty"`
	TransFatServing         float64 `json:"trans-fat_serving,omitempty"`
	TransFatValue           float64 `json:"trans-fat_value,omitempty"`
	VitaminA                float64 `json:"vitamin-a,omitempty"`
	VitaminA100G            float64 `json:"vitamin-a_100g,omitempty"`
	VitaminAValue           float64 `json:"vitamin-a_value,omitempty"`
	VitaminAServing         float64 `json:"vitamin-a_serving,omitempty"`
	VitaminAUnit            string  `json:"vitamin-a_unit,omitempty"`
	VitaminALabel           string  `json:"vitamin-a_label,omitempty"`
	VitaminCValue           float64 `json:"vitamin-c_value,omitempty"`
	VitaminCUnit            string  `json:"vitamin-c_unit,omitempty"`
	VitaminC100G            float64 `json:"vitamin-c_100g,omitempty"`
	VitaminC                float64 `json:"vitamin-c,omitempty"`
	VitaminCServing         float64 `json:"vitamin-c_serving,omitempty"`
	VitaminCLabel           string  `json:"vitamin-c_label,omitempty"`
	Proteins100G            float64 `json:"proteins_100g,omitempty"`
	ProteinsServing         float64 `json:"proteins_serving,omitempty"`
	ProteinsValue           float64 `json:"proteins_value,omitempty"`
	ProteinsUnit            string  `json:"proteins_unit,omitempty"`
	Proteins                float64 `json:"proteins,omitempty"`
	PolyunsaturatedFat100G  float64 `json:"polyunsaturated-fat_100g,omitempty"`
	Potassium100G           float64 `json:"potassium_100g,omitempty"`
	Sodium                  float64 `json:"sodium,omitempty"`
	SodiumServing           float64 `json:"sodium_serving,omitempty"`
	SodiumValue             float64 `json:"sodium_value,omitempty"`
	Sodium100G              float64 `json:"sodium_100g,omitempty"`
	SodiumUnit              string  `json:"sodium_unit,omitempty"`
	CarbohydratesUnit       string  `json:"carbohydrates_unit,omitempty"`
	CarbohydratesValue      float64 `json:"carbohydrates_value,omitempty"`
	Carbohydrates100G       float64 `json:"carbohydrates_100g,omitempty"`
	Carbohydrates           string  `json:"carbohydrates,omitempty"`
	CarbohydratesServing    float64 `json:"carbohydrates_serving,omitempty"`
	AlcoholValue            float64 `json:"alcohol_value,omitempty"`
	AlcoholServing          float64 `json:"alcohol_serving,omitempty"`
	AlcoholUnit             string  `json:"alcohol_unit,omitempty"`
	Alcohol100G             float64 `json:"alcohol_100g,omitempty"`
	Alcohol                 float64 `json:"alcohol,omitempty"`
	MonounsaturatedFat100G  float64 `json:"monounsaturated-fat_100g,omitempty"`
	NovaGroup               float64 `json:"nova-group,omitempty"`
	NovaGroupServing        float64 `json:"nova-group_serving,omitempty"`
	NovaGroup100G           float64 `json:"nova-group_100g,omitempty"`
	Energy                  float64 `json:"energy,omitempty"`
	EnergyServing           float64 `json:"energy_serving,omitempty"`
	EnergyKcalServing       float64 `json:"energy-kcal_serving,omitempty"`
	EnergyKcal              float64 `json:"energy-kcal,omitempty"`
	Energy100G              float64 `json:"energy_100g,omitempty"`
	EnergyUnit              string  `json:"energy_unit,omitempty"`
	EnergyKcalValue         float64 `json:"energy-kcal_value,omitempty"`
	EnergyKcalUnit          string  `json:"energy-kcal_unit,omitempty"`
	EnergyKcal100G          float64 `json:"energy-kcal_100g,omitempty"`
	EnergyValue             float64 `json:"energy_value,omitempty"`
	NutritionScoreUk100G    float64 `json:"nutrition-score-uk_100g,omitempty"`
	NutritionScoreFrServing float64 `json:"nutrition-score-fr_serving,omitempty"`
	NutritionScoreFr        float64 `json:"nutrition-score-fr,omitempty"`
	NutritionScoreFr100G    float64 `json:"nutrition-score-fr_100g,omitempty"`
	NutritionScoreUkServing float64 `json:"nutrition-score-uk_serving,omitempty"`
	NutritionScoreUk        float64 `json:"nutrition-score-uk,omitempty"`
	Fiber                   float64 `json:"fiber,omitempty"`
	Fiber100G               float64 `json:"fiber_100g,omitempty"`
	FiberValue              float64 `json:"fiber_value,omitempty"`
	FiberServing            float64 `json:"fiber_serving,omitempty"`
	FiberUnit               string  `json:"fiber_unit,omitempty"`
}
