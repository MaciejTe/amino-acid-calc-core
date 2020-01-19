package ingredient

import (
	"github.com/MaciejTe/amino-acid-calc/pkg/usda"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	food string
)

// Search searches for possible ingredients in USDA database.
func Search() *cobra.Command {
	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search for food in USDA database",
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
			ingredientSearch()
		},
	}
	searchCmd.PersistentFlags().StringVarP(&food, "food", "f", "", "Food to search for")
	return searchCmd
}


func ingredientSearch() {
	log.Debugf("Running command food search with parameter %v", food)
	config := viper.GetStringMapString("core")
	if config["usda_api_key"] == "" {
		log.Error("USDA API key configuration missing")
		return
	}
	client := usda.NewClient(config["usda_api_key"])
	resp, err := client.FoodSearch(food)
	if err != nil {
		log.Error("Food search request error: ", err)
		return
	}
	log.Debug("USDA response code: ", resp.StatusCode())
	spew.Dump(resp)
}
