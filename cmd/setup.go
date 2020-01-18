package cmd

import (
	"github.com/MaciejTe/amino-acid-calc/cmd/calculate"
	"github.com/MaciejTe/amino-acid-calc/cmd/ingredient"
	"github.com/spf13/viper"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// nolint: gochecknoglobals
	cfgFile string
)

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

// Execute initializes all CLI commands inside application.
func Execute() int {
	var rootCmd = &cobra.Command{
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

	var ingredientCmd = &cobra.Command {
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

	var calculateCmd = &cobra.Command{
		Use:   "calculate",
		Short: "Calculation section",
		Long:  "Commands related to calculating macro and microelements in recipes",
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
	rootCmd.AddCommand(calculateCmd)

	ingredientCmd.AddCommand(ingredient.Search())
	ingredientCmd.AddCommand(ingredient.Details())

	calculateCmd.AddCommand(calculate.Recipe())

	cobra.OnInitialize(initConfig)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}
