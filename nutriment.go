// Copyright © 2019 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

import (
	"fmt"
	"strconv"
)

type StrFloat float64

func (s *StrFloat) UnmarshalJSON(data []byte) error {
	if string(data) == "0" {
		*s = -1
		return nil
	}

	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("malformed data")
	}

	v, err := strconv.ParseFloat(string(data[1:len(data)-1]), 64)
	if err != nil {
		return err
	}
	*s = StrFloat(v)
	return nil
}

// Nutriment is the collection of all retruned nutriment associated with the product
type Nutriment struct {
	Salt                    StrFloat `json:"salt,omitempty"`
	Salt100G                StrFloat `json:"salt_100g,omitempty"`
	SaltValue               StrFloat `json:"salt_value,omitempty"`
	SaltServing             StrFloat `json:"salt_serving,omitempty"`
	SaltUnit                string   `json:"salt_unit,omitempty"`
	Sugars100G              StrFloat `json:"sugars_100g,omitempty"`
	Sugars                  StrFloat `json:"sugars,omitempty"`
	SugarsUnit              string   `json:"sugars_unit,omitempty"`
	SugarsServing           StrFloat `json:"sugars_serving,omitempty"`
	SugarsValue             StrFloat `json:"sugars_value,omitempty"`
	Iron                    StrFloat `json:"iron,omitempty"`
	IronValue               StrFloat `json:"iron_value,omitempty"`
	IronLabel               string   `json:"iron_label,omitempty"`
	IronUnit                string   `json:"iron_unit,omitempty"`
	Iron100G                StrFloat `json:"iron_100g,omitempty"`
	IronServing             StrFloat `json:"iron_serving,omitempty"`
	CalciumUnit             string   `json:"calcium_unit,omitempty"`
	CalciumServing          StrFloat `json:"calcium_serving,omitempty"`
	Calcium                 StrFloat `json:"calcium,omitempty"`
	CalciumValue            StrFloat `json:"calcium_value,omitempty"`
	CalciumLabel            string   `json:"calcium_label,omitempty"`
	Calcium100G             StrFloat `json:"calcium_100g,omitempty"`
	Cholesterol100G         StrFloat `json:"cholesterol_100g,omitempty"`
	SaturatedFat100G        StrFloat `json:"saturated-fat_100g,omitempty"`
	SaturatedFatServing     StrFloat `json:"saturated-fat_serving,omitempty"`
	SaturatedFat            StrFloat `json:"saturated-fat,omitempty"`
	SaturatedFatValue       StrFloat `json:"saturated-fat_value,omitempty"`
	SaturatedFatUnit        string   `json:"saturated-fat_unit,omitempty"`
	Fat100G                 StrFloat `json:"fat_100g,omitempty"`
	FatServing              StrFloat `json:"fat_serving,omitempty"`
	FatValue                StrFloat `json:"fat_value,omitempty"`
	FatUnit                 string   `json:"fat_unit,omitempty"`
	Fat                     StrFloat `json:"fat,omitempty"`
	TransFatLabel           string   `json:"trans-fat_label,omitempty"`
	TransFatUnit            string   `json:"trans-fat_unit,omitempty"`
	TransFat                StrFloat `json:"trans-fat,omitempty"`
	TransFat100G            StrFloat `json:"trans-fat_100g,omitempty"`
	TransFatServing         StrFloat `json:"trans-fat_serving,omitempty"`
	TransFatValue           StrFloat `json:"trans-fat_value,omitempty"`
	VitaminA                StrFloat `json:"vitamin-a,omitempty"`
	VitaminA100G            StrFloat `json:"vitamin-a_100g,omitempty"`
	VitaminAValue           StrFloat `json:"vitamin-a_value,omitempty"`
	VitaminAServing         StrFloat `json:"vitamin-a_serving,omitempty"`
	VitaminAUnit            string   `json:"vitamin-a_unit,omitempty"`
	VitaminALabel           string   `json:"vitamin-a_label,omitempty"`
	VitaminCValue           StrFloat `json:"vitamin-c_value,omitempty"`
	VitaminCUnit            string   `json:"vitamin-c_unit,omitempty"`
	VitaminC100G            StrFloat `json:"vitamin-c_100g,omitempty"`
	VitaminC                StrFloat `json:"vitamin-c,omitempty"`
	VitaminCServing         StrFloat `json:"vitamin-c_serving,omitempty"`
	VitaminCLabel           string   `json:"vitamin-c_label,omitempty"`
	Proteins100G            StrFloat `json:"proteins_100g,omitempty"`
	ProteinsServing         StrFloat `json:"proteins_serving,omitempty"`
	ProteinsValue           StrFloat `json:"proteins_value,omitempty"`
	ProteinsUnit            string   `json:"proteins_unit,omitempty"`
	Proteins                StrFloat `json:"proteins,omitempty"`
	PolyunsaturatedFat100G  StrFloat `json:"polyunsaturated-fat_100g,omitempty"`
	Potassium100G           StrFloat `json:"potassium_100g,omitempty"`
	Sodium                  StrFloat `json:"sodium,omitempty"`
	SodiumServing           StrFloat `json:"sodium_serving,omitempty"`
	SodiumValue             StrFloat `json:"sodium_value,omitempty"`
	Sodium100G              StrFloat `json:"sodium_100g,omitempty"`
	SodiumUnit              string   `json:"sodium_unit,omitempty"`
	CarbohydratesUnit       string   `json:"carbohydrates_unit,omitempty"`
	CarbohydratesValue      StrFloat `json:"carbohydrates_value,omitempty"`
	Carbohydrates100G       StrFloat `json:"carbohydrates_100g,omitempty"`
	Carbohydrates           StrFloat `json:"carbohydrates,omitempty"`
	CarbohydratesServing    StrFloat `json:"carbohydrates_serving,omitempty"`
	AlcoholValue            StrFloat `json:"alcohol_value,omitempty"`
	AlcoholServing          StrFloat `json:"alcohol_serving,omitempty"`
	AlcoholUnit             string   `json:"alcohol_unit,omitempty"`
	Alcohol100G             StrFloat `json:"alcohol_100g,omitempty"`
	Alcohol                 StrFloat `json:"alcohol,omitempty"`
	MonounsaturatedFat100G  StrFloat `json:"monounsaturated-fat_100g,omitempty"`
	NovaGroup               StrFloat `json:"nova-group,omitempty"`
	NovaGroupServing        StrFloat `json:"nova-group_serving,omitempty"`
	NovaGroup100G           StrFloat `json:"nova-group_100g,omitempty"`
	Energy                  StrFloat `json:"energy,omitempty"`
	EnergyServing           StrFloat `json:"energy_serving,omitempty"`
	EnergyKcalServing       StrFloat `json:"energy-kcal_serving,omitempty"`
	EnergyKcal              StrFloat `json:"energy-kcal,omitempty"`
	Energy100G              StrFloat `json:"energy_100g,omitempty"`
	EnergyUnit              string   `json:"energy_unit,omitempty"`
	EnergyKcalValue         StrFloat `json:"energy-kcal_value,omitempty"`
	EnergyKcalUnit          string   `json:"energy-kcal_unit,omitempty"`
	EnergyKcal100G          StrFloat `json:"energy-kcal_100g,omitempty"`
	EnergyValue             StrFloat `json:"energy_value,omitempty"`
	NutritionScoreUk100G    StrFloat `json:"nutrition-score-uk_100g,omitempty"`
	NutritionScoreFrServing StrFloat `json:"nutrition-score-fr_serving,omitempty"`
	NutritionScoreFr        StrFloat `json:"nutrition-score-fr,omitempty"`
	NutritionScoreFr100G    StrFloat `json:"nutrition-score-fr_100g,omitempty"`
	NutritionScoreUkServing StrFloat `json:"nutrition-score-uk_serving,omitempty"`
	NutritionScoreUk        StrFloat `json:"nutrition-score-uk,omitempty"`
	Fiber                   StrFloat `json:"fiber,omitempty"`
	Fiber100G               StrFloat `json:"fiber_100g,omitempty"`
	FiberValue              StrFloat `json:"fiber_value,omitempty"`
	FiberServing            StrFloat `json:"fiber_serving,omitempty"`
	FiberUnit               string   `json:"fiber_unit,omitempty"`
}
