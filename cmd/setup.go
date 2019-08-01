package cmd

import (
	_ "github.com/davecgh/go-spew/spew"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// nolint: gochecknoglobals
	cfgFile string
)

func setupCommands(rootCmd *cobra.Command) {
	Serve(rootCmd)
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
		log.Debug("Loading config file from: config.yaml",)
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

func Execute() {
	var rootCmd = &cobra.Command {
		Use:   "application",
		Short: "Golang CLI/REST API application",
		Long:  "Sample Golang application",
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

	setupCommands(rootCmd)

	cobra.OnInitialize(initConfig)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
