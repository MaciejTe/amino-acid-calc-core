package ingredient

import (
	"github.com/MaciejTe/amino-acid-calc/pkg"
	"github.com/davecgh/go-spew/spew"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/MaciejTe/amino-acid-calc/pkg/usda"
)

var (
	ingredientID string
)

func IngredientDetails(ingredientCmd *cobra.Command) {
	detailsCmd := &cobra.Command{
		Use:   "details",
		Short: "Search for food details in USDA database",
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
			log.Debug("Running command food details with parameter ", ingredientID)
			config := viper.GetStringMapString("core")
			if config["usda_api_key"] == "" {
				log.Error("USDA API key configuration missing")
				return
			}
			client := usda.NewClient(config["usda_api_key"])
			resp, err := client.FoodDetails(ingredientID)
			if err != nil {
				log.Error("Ingredient search request error: ", err)
				return
			}
			log.Debug("USDA response: ", resp.StatusCode())
			foodDetails, err := pkg.NewIngredient(ingredientID, resp.Body())
			if err != nil {
				log.Error("Failed to convert USDA response into Ingredient struct: ", err)
			}
			foodDetailsStr := spew.Sdump(foodDetails)
			log.Info("Returned food information: ", foodDetailsStr)
		},
	}
	detailsCmd.PersistentFlags().StringVarP(&ingredientID, "fid", "i", "", "Ingredient ID to search for")
	ingredientCmd.AddCommand(detailsCmd)
}
