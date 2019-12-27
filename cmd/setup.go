package cmd

import (
	"github.com/MaciejTe/amino-acid-calc/cmd/ingredient"
	_ "github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// nolint: gochecknoglobals
	cfgFile string
)

func setupIngredientCommands(ingredientCmd *cobra.Command) {
	ingredient.IngredientSearch(ingredientCmd)
	ingredient.IngredientDetails(ingredientCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		log.Debug("Loading config file from: config.yaml")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath(dir)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if cfgFile != "" {
			log.Warning("Config specified but unable to read it, using defaults")
		}
	}
}

func Execute() int {
	var rootCmd = &cobra.Command {
		Use:   "amino-acid-calc-core",
		Short: "Amino acids calculator application",
		Long:  "Amino acids calculator application",
		Run: func(cmd *cobra.Command, args []string) {
			// print help and quit
			if len(args) == 0 {
				err := cmd.Help()
				if err != nil {
					log.Error(err)
					os.Exit(1)
				}
				os.Exit(0)
			}
		},
	}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	var ingredientCmd = &cobra.Command{
		Use:   "ingredient",
		Short: "Ingredient related commands",
		Long:  "Ingredient related commands",
		Run: func(cmd *cobra.Command, args []string) {
			// print help and quit
			if len(args) == 0 {
				err := cmd.Help()
				if err != nil {
					log.Error(err)
					os.Exit(1)
				}
				os.Exit(0)
			}
		},
	}
	rootCmd.AddCommand(ingredientCmd)

	setupIngredientCommands(ingredientCmd)

	cobra.OnInitialize(initConfig)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}
