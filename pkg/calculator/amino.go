package calculator

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
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

// NewIngredient converts USDA REST API JSON response and returns Ingredient structure.
func NewIngredient(ingredientID string, usdaResponseBody []byte) (foodDetails *Ingredient, err error) {
	if ingredientID == "" {
		errMsg := "NewIngredient received empty ingredientID string"
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	foodDetails = &Ingredient{}
	var responseMapTemplate interface{}
	log.Debugf("Launching NewIngredient for ingredient %v...", ingredientID)
	err = json.Unmarshal(usdaResponseBody, &responseMapTemplate)
	if err != nil {
		log.Error("Cannot unmarshal JSON response: ", err)
		return
	}
	responseMap := responseMapTemplate.(map[string]interface{})
	foodDetails.Description = cast.ToString(responseMap["description"])
	foodDetails.ID = ingredientID
	for _, nutrientMap := range responseMap["foodNutrients"].([]interface{}) {
		for key, nutrientNameMap := range nutrientMap.(map[string]interface{}) {
			if key == "nutrient" {
				nutrientNameMapCasted := nutrientNameMap.(map[string]interface{})
				nutrientMapCasted := nutrientMap.(map[string]interface{})
				switch nutrientNameMapCasted["name"] {
				case "Carbohydrate, by summation":
					foodDetails.Carbohydrates = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Protein":
					foodDetails.Protein = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Energy":
					foodDetails.Kcal = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Total lipid (fat)":
					foodDetails.Fat = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Alanine":
					foodDetails.AminoAcids.Alanine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Isoleucine":
					foodDetails.AminoAcids.Isoleucine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Leucine":
					foodDetails.AminoAcids.Leucine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Valine":
					foodDetails.AminoAcids.Valine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Phenylalanine":
					foodDetails.AminoAcids.Phenylalanine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Tryptophan":
					foodDetails.AminoAcids.Tryptophan = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Tyrosine":
					foodDetails.AminoAcids.Tyrosine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Asparagine":
					foodDetails.AminoAcids.Asparagine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Cysteine":
					foodDetails.AminoAcids.Cysteine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Glutamine":
					foodDetails.AminoAcids.Glutamine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Methionine":
					foodDetails.AminoAcids.Methionine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Serine":
					foodDetails.AminoAcids.Serine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Threonine":
					foodDetails.AminoAcids.Threonine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Arginine":
					foodDetails.AminoAcids.Arginine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Histidine":
					foodDetails.AminoAcids.Histidine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Lysine":
					foodDetails.AminoAcids.Lysine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Aspartic acid":
					foodDetails.AminoAcids.AsparticAcid = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Glutamic acid":
					foodDetails.AminoAcids.GlutamicAcid = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Glycine":
					foodDetails.AminoAcids.Glycine = cast.ToFloat32(nutrientMapCasted["amount"])
				case "Proline":
					foodDetails.AminoAcids.Proline = cast.ToFloat32(nutrientMapCasted["amount"])
				}
			}
		}
	}
	log.Debugf("Launching NewIngredient for ingredient %v... DONE", ingredientID)
	return
}
