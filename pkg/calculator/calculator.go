package calculator

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/MaciejTe/amino-acid-calc/db"
	"github.com/MaciejTe/amino-acid-calc/pkg/usda"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type FoodSearchResponseBody struct {
	FoodSearchCriteria map[string]interface{} `json:"foodSearchCriteria"`
	TotalHits int `json:"totalHits"`
	CurrentPage int `json:"currentPage"`
	TotalPages int `json:"totalPages"`
	Foods []map[string]interface{} `json:"foods"`
}

type Calculator interface {
	GetRecipeNutritionFacts()
}

// Recipe structure
type Recipe struct {
	Id int
	Name string
	Author string
	Description string
	Ingredients []byte
	Instructions []byte
	Servings string
	Link string
	Duration string
	Category string
	Calculated bool
	Success bool
	NutritionFacts sql.NullInt64

	nutritionFacts *NutritionFacts
	usdaClient     usda.USDAClient
	dbClient       *sql.DB
}

func NewRecipe() (*Recipe, error) {
	// create USDA DB client
	config := viper.GetStringMapString("core")
	if config["usda_api_key"] == "" {
		errMsg := "NewRecipe: USDA API key configuration missing"
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	client := usda.NewClient(config["usda_api_key"])

	// create internal DB client
	dbClient, err := db.NewDbClient()
	if err != nil {
		return nil, err
	}

	return &Recipe{
		usdaClient: *client,
		dbClient:	dbClient,
	}, nil
}

// nutritionFacts structure contains list of calculated ingredient components and sum of them
type NutritionFacts struct {
	NutritionSlice []*Ingredient
	NutritionFactsSum *NutritionFactsSum
}

func NewNutritionFacts(nutritionSlice []*Ingredient) *NutritionFacts {
	nutritionSum := NewNutritionFactsSum(nutritionSlice)
	nutritionFacts := NutritionFacts{
		NutritionSlice:    nutritionSlice,
		NutritionFactsSum: nutritionSum,
	}
	return &nutritionFacts
}

// NewNutritionFactsSum sums up each ingredient nutrition facts
func NewNutritionFactsSum(nutritionSlice []*Ingredient) *NutritionFactsSum {
	sum := NutritionFactsSum{}
	for _, ingredient := range nutritionSlice {
		if ingredient == nil {
			// ingredient was not found during food search
			continue
		}
		sum.Carbohydrates += ingredient.Carbohydrates
		sum.Kcal += ingredient.Kcal
		sum.Protein += ingredient.Protein
		sum.Fat += ingredient.Fat

		sum.aminoAcids.Alanine += ingredient.AminoAcids.Alanine
		sum.aminoAcids.Isoleucine += ingredient.AminoAcids.Isoleucine
		sum.aminoAcids.Leucine += ingredient.AminoAcids.Leucine
		sum.aminoAcids.Valine += ingredient.AminoAcids.Valine

		sum.aminoAcids.Phenylalanine += ingredient.AminoAcids.Alanine
		sum.aminoAcids.Tryptophan += ingredient.AminoAcids.Alanine
		sum.aminoAcids.Tyrosine += ingredient.AminoAcids.Alanine

		sum.aminoAcids.Asparagine += ingredient.AminoAcids.Asparagine
		sum.aminoAcids.Cysteine += ingredient.AminoAcids.Cysteine
		sum.aminoAcids.Glutamine += ingredient.AminoAcids.Glutamine
		sum.aminoAcids.Methionine += ingredient.AminoAcids.Methionine
		sum.aminoAcids.Serine += ingredient.AminoAcids.Serine
		sum.aminoAcids.Threonine += ingredient.AminoAcids.Threonine

		sum.aminoAcids.Arginine += ingredient.AminoAcids.Arginine
		sum.aminoAcids.Histidine += ingredient.AminoAcids.Histidine
		sum.aminoAcids.Lysine += ingredient.AminoAcids.Lysine

		sum.aminoAcids.AsparticAcid += ingredient.AminoAcids.AsparticAcid
		sum.aminoAcids.GlutamicAcid += ingredient.AminoAcids.GlutamicAcid

		sum.aminoAcids.Glycine += ingredient.AminoAcids.Glycine
		sum.aminoAcids.Proline += ingredient.AminoAcids.Proline

	}
	return &sum
}

type NutritionFactsSum struct {
	//RecipeName string
	Carbohydrates float32
	Protein       float32
	Fat           float32
	Kcal          float32
	AminoAcids    sql.NullInt64

	aminoAcids AminoAcids
}

func (r *Recipe) GetRecipeNutritionFacts() (*NutritionFacts, error) {
	ingredientsList, err := r.convertListOfBytesToListOfMaps()
	if err != nil {
		return nil, err
	}
	log.Debugf("Ingredients list length: %v", len(ingredientsList))
	var nutritionFacts []*Ingredient
	for _, ingredient := range ingredientsList {
		foodId, err := r.IngredientSearch(ingredient["name"])
		if err != nil {
			r.markRecipeAsCalculated(false)
			return nil, err
		}
		// omitting error handling here, not getting all ingredient details is allowed
		ingredient, _ := r.IngredientDetails(foodId)
		nutritionFacts = append(nutritionFacts, ingredient)
	}
	r.nutritionFacts = NewNutritionFacts(nutritionFacts)
	if err := r.saveAminoAcidsToDB(); err != nil {
		return nil, err
	}
	if err := r.saveNutritionFactsToDB(); err != nil {
		return nil, err
	}
	r.markRecipeAsCalculated(true)

	return r.nutritionFacts, nil
}

func (r *Recipe) saveAminoAcidsToDB() error {
	lastInsertId := 0
	insertStatement := "INSERT INTO amino_acids(alanine, isoleucine, leucine, valine, phenylalanine, tryptophan, tyrosine, asparagine, cysteine, " +
		"glutamine, methionine, serine, threonine, arginine, histidine, lysine, aspartic_acid, glutamic_acid, glycine, proline) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20) RETURNING id;"
	aminoAcids := r.nutritionFacts.NutritionFactsSum.aminoAcids
	err := r.dbClient.QueryRow(insertStatement, aminoAcids.Alanine, aminoAcids.Isoleucine, aminoAcids.Leucine,
		aminoAcids.Valine, aminoAcids.Phenylalanine, aminoAcids.Tryptophan, aminoAcids.Tyrosine, aminoAcids.Asparagine,
		aminoAcids.Cysteine, aminoAcids.Glutamine, aminoAcids.Methionine, aminoAcids.Serine, aminoAcids.Threonine,
		aminoAcids.Arginine, aminoAcids.Histidine, aminoAcids.Lysine, aminoAcids.AsparticAcid, aminoAcids.GlutamicAcid,
		aminoAcids.Glycine, aminoAcids.Proline).Scan(&lastInsertId)
	if err != nil {
		log.Errorf("Failed to insert amino acids to database! Details: %v", err)
		return err
	}
	r.nutritionFacts.NutritionFactsSum.AminoAcids.Int64 = int64(lastInsertId)
	log.Debugf("Saved amino acids to database for recipe: %v", r.Link)
	return err
}

func (r *Recipe) saveNutritionFactsToDB() error {
	lastInsertId := 0
	insertStatement := "INSERT INTO nutrition_facts(amino_acids, carbohydrates, protein, fat, kcal) " +
		"VALUES ($1, $2, $3, $4, $5) RETURNING id;"
	err := r.dbClient.QueryRow(insertStatement, r.nutritionFacts.NutritionFactsSum.AminoAcids.Int64,
		r.nutritionFacts.NutritionFactsSum.Carbohydrates, r.nutritionFacts.NutritionFactsSum.Protein,
		r.nutritionFacts.NutritionFactsSum.Fat, r.nutritionFacts.NutritionFactsSum.Kcal).Scan(&lastInsertId)
	if err != nil {
		log.Errorf("Failed to insert nutrition facts to database! Details: %v", err)
		return err
	}
	updateStatement := "UPDATE recipes SET nutrition_facts = $1 WHERE link = $2;"
	_, err = r.dbClient.Exec(updateStatement, lastInsertId, r.Link)
	if err != nil {
		log.Errorf("Failed to update recipe nutrition facts in database! Details: %v", err)
		return err
	}
	log.Debugf("Saved nutrition facts to database for recipe: %v", r.Link)

	return nil
}

func (r *Recipe) markRecipeAsCalculated(success bool) error {
	sqlStatement := "UPDATE recipes SET calculated = true, success = $1 WHERE link = $2;"
	_, err := r.dbClient.Exec(sqlStatement, success, r.Link)
	if err != nil {
		return err
	}
	log.Debugf("Marked recipe with link %v as calculated with success value %v", r.Link, success)

	r.Calculated = true
	r.Success = success
	return nil
}

func (r *Recipe) IngredientSearch(food string) (ingredientId string, err error) {
	resp, err := r.usdaClient.FoodSearch(food)
	if err != nil {
		log.Error("Food search request error: ", err)
		return
	}
	if resp.IsError() {
		return
	}
	log.Debug("USDA response code: ", resp.StatusCode())
	foodResponseStructure := FoodSearchResponseBody{}
	err = json.Unmarshal(resp.Body(), &foodResponseStructure)
	if err != nil {
		log.Errorf("Failed to unmarshal USDA search response! Details: %v", err)
		return
	}
	// TODO: unfortunately I had to neglect accuracy of found ingredients - it's hard to choose correct word to search for
	// TODO: in situation where ingredient is described as "thinly sliced red onion" or something similar, instead of "red onion";
	// TODO: for now first found food is chosen; in future better behaviour will be designed
	if len(foodResponseStructure.Foods) == 0 {
		errMsg := "No food with given name was found in USDA database: " + food
		log.Warnf(errMsg)
		return "", errors.New(errMsg)
	}
	log.Infof("Searching for food: %v | Found food: %v", food,foodResponseStructure.Foods[0]["description"])
	fdcId, err := cast.ToStringE(foodResponseStructure.Foods[0]["fdcId"])
	if err != nil {
		return "", err
	}
	return fdcId, nil
}

func (r *Recipe) IngredientDetails(foodId string) (*Ingredient, error) {
	resp, err := r.usdaClient.FoodDetails(foodId)
	if err != nil {
		return nil, err
	}
	ingredient, err := NewIngredient(foodId, resp.Body())
	return ingredient, err
}

func (r *Recipe) convertListOfBytesToListOfMaps() ([]map[string]string, error) {
	var ingredientsListOfStrings []string
	var ingredientsListMap []map[string]string

	err := json.Unmarshal(r.Ingredients, &ingredientsListOfStrings)
	if err != nil {
		log.Fatalf("Failed to convert ingredients to slice. Details: %v", err)
	}
	for _, ingredientString := range ingredientsListOfStrings {
		ingredientMap := map[string]string{}
		err := json.Unmarshal([]byte(ingredientString), &ingredientMap)
		if err != nil {
			log.Fatalf("Failed to unmarshal list of ingredients! Details: %v", err)
		}
		ingredientsListMap = append(ingredientsListMap, ingredientMap)
	}
	return ingredientsListMap, err
}
