package calculator

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
)

// Ingredient structure holds information about all major food data.
type Ingredient struct {
	ID            string
	Description   string
	AminoAcids    AminoAcids
	Carbohydrates float32
	Protein       float32
	Fat           float32
	Kcal          float32
}

// aminoAcids structure holds all necessary amino acids.
type AminoAcids struct {
	// Aliphatic Amino Acids with Hydrophobic Side Chain
	Alanine    float32
	Isoleucine float32
	Leucine    float32
	Valine     float32

	// Aromatic Amino Acids with Hydrophobic Side Chain
	Phenylalanine float32
	Tryptophan    float32
	Tyrosine      float32

	// Amino Acids with Neutral Side Chain
	Asparagine float32
	Cysteine   float32
	Glutamine  float32
	Methionine float32
	Serine     float32
	Threonine  float32

	// Amino Acids with Positive Charged Side Chain
	Arginine  float32
	Histidine float32
	Lysine    float32

	// Amino Acids with Negative Charged Side Chain
	AsparticAcid float32
	GlutamicAcid float32

	// Unique Amino Acids
	Glycine float32
	Proline float32
}

type Nutrient struct {
	NutrientData struct {
		Name string `json:"name"`
		UnitName string `json:"unitName"`
	} `json:"nutrient"`
	Amount float32 `json:"amount"`
}

type UsdaDetailResponse struct {
	Description string `json:"description"`
	FoodNutrients []Nutrient `json:"foodNutrients"`
}

// NewIngredient converts USDA REST API JSON response and returns Ingredient structure.
func NewIngredient(ingredientID string, usdaResponseBody []byte) (foodDetails *Ingredient, err error) {
	if ingredientID == "" {
		errMsg := "NewIngredient received empty ingredientID string"
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	foodDetails = &Ingredient{}
	usdaDetailResponse := UsdaDetailResponse{}
	log.Debugf("Launching NewIngredient for ingredient %v...", ingredientID)
	err = json.Unmarshal(usdaResponseBody, &usdaDetailResponse)
	if err != nil {
		errMsg := "cannot unmarshal JSON response: " + err.Error()
		log.Errorf(errMsg)
		return nil, errors.New(errMsg)
	}
	foodDetails.Description = usdaDetailResponse.Description
	foodDetails.ID = ingredientID
	for _, nutrient := range usdaDetailResponse.FoodNutrients {
		switch nutrient.NutrientData.Name {
		case "Carbohydrate, by summation":
			foodDetails.Carbohydrates = nutrient.Amount
		case "Protein":
			foodDetails.Protein = nutrient.Amount
		case "Energy":
			if nutrient.NutrientData.UnitName == "kcal" {
				foodDetails.Kcal = nutrient.Amount
			}
		case "Total lipid (fat)":
			foodDetails.Fat = nutrient.Amount
		case "Alanine":
			foodDetails.AminoAcids.Alanine = nutrient.Amount
		case "Isoleucine":
			foodDetails.AminoAcids.Isoleucine = nutrient.Amount
		case "Leucine":
			foodDetails.AminoAcids.Leucine = nutrient.Amount
		case "Valine":
			foodDetails.AminoAcids.Valine = nutrient.Amount
		case "Phenylalanine":
			foodDetails.AminoAcids.Phenylalanine = nutrient.Amount
		case "Tryptophan":
			foodDetails.AminoAcids.Tryptophan = nutrient.Amount
		case "Tyrosine":
			foodDetails.AminoAcids.Tyrosine = nutrient.Amount
		case "Asparagine":
			foodDetails.AminoAcids.Asparagine = nutrient.Amount
		case "Cysteine":
			foodDetails.AminoAcids.Cysteine = nutrient.Amount
		case "Glutamine":
			foodDetails.AminoAcids.Glutamine = nutrient.Amount
		case "Methionine":
			foodDetails.AminoAcids.Methionine = nutrient.Amount
		case "Serine":
			foodDetails.AminoAcids.Serine = nutrient.Amount
		case "Threonine":
			foodDetails.AminoAcids.Threonine = nutrient.Amount
		case "Arginine":
			foodDetails.AminoAcids.Arginine = nutrient.Amount
		case "Histidine":
			foodDetails.AminoAcids.Histidine = nutrient.Amount
		case "Lysine":
			foodDetails.AminoAcids.Lysine = nutrient.Amount
		case "Aspartic acid":
			foodDetails.AminoAcids.AsparticAcid = nutrient.Amount
		case "Glutamic acid":
			foodDetails.AminoAcids.GlutamicAcid = nutrient.Amount
		case "Glycine":
			foodDetails.AminoAcids.Glycine = nutrient.Amount
		case "Proline":
			foodDetails.AminoAcids.Proline = nutrient.Amount
		}
	}
	log.Debugf("Launching NewIngredient for ingredient %v... DONE", ingredientID)
	return
}
