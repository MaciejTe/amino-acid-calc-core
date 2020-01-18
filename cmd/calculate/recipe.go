package calculate

import (
	"github.com/MaciejTe/amino-acid-calc/db"
	"github.com/MaciejTe/amino-acid-calc/pkg/calculator"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

var (
	ingredientID string
)

// Recipe gets uncalculated recipe from internal DB and calulates its macro and microelements with use of USDA database.
func Recipe() *cobra.Command {
	calcRecipeCmd := &cobra.Command{
		Use:   "recipe",
		Short: "Calculate macro and microelements in recipe with use of USDA database",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			log.SetLevel(log.DebugLevel)
			log.SetFormatter(&log.JSONFormatter{
				FieldMap: log.FieldMap{
					log.FieldKeyTime:  "timestamp",
					log.FieldKeyLevel: "loglevel",
					log.FieldKeyMsg:   "message",
				},
				TimestampFormat: time.RFC3339,
			})
		},
		Run: func(cmd *cobra.Command, args []string) {
			log.Debugf("Running command calculate recipe with parameter ", ingredientID)
			dbClient, err := db.NewDbClient()
			if err != nil {
				log.Fatalf("Cannot create DB client! Details: %v", err)
			}
			selectStatement := "SELECT * FROM recipes WHERE calculated = false LIMIT 1;"
			recipe, err := calculator.NewRecipe()
			if err != nil {
				log.Errorf("Failed to create recipe! Details: %v", err)
			}
			err = dbClient.QueryRow(selectStatement).Scan(&recipe.Id, &recipe.Name, &recipe.Author, &recipe.Description,
				&recipe.Ingredients, &recipe.Instructions, &recipe.Servings, &recipe.Link, &recipe.Duration,
				&recipe.Category, &recipe.Calculated, &recipe.Success, &recipe.NutritionFacts)
			if err != nil {
				log.Fatalf("Failed to scan returned query into Recipe structure! Details: %v", err)
			}
			log.Debug("Returned row successfully scanned into Recipe structure")
			nutritionFacts, err := recipe.GetRecipeNutritionFacts()
			spew.Dump(nutritionFacts, err, "::::::")
		},
	}
	//detailsCmd.PersistentFlags().StringVarP(&ingredientID, "fid", "i", "", "Ingredient ID to search for")
	return calcRecipeCmd
}
