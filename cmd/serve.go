package cmd

import (
	"time"

	"github.com/MaciejTe/go-project-template/pkg/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ApiPort string
	defaultPortVal = ":8080"
)

func Serve(rootCmd *cobra.Command) {
	serveCmd := &cobra.Command {
		Use:   "serve",
		Short: "Start REST API server",
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
			router := api.SetupRouter()
			log.Debug("Running command serve...")
			apiMap := viper.GetStringMapString("api")
			if len(apiMap["port"]) != 0  && ApiPort == defaultPortVal{
				ApiPort = apiMap["port"]
			}
			log.Debug("Running REST API on port ", ApiPort)

			router.Run(ApiPort)
		},
	}
	serveCmd.PersistentFlags().StringVarP(&ApiPort, "port", "p", defaultPortVal, "API port")
	//viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	//viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))
	rootCmd.AddCommand(serveCmd)
}
