package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/golang-rennes/livecoding/config"
	"github.com/golang-rennes/livecoding/models"
	"github.com/golang-rennes/livecoding/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var mainCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Long:  "The version of user-list.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version 0.0.1")
	},
}

func init() {
	mainCmd.AddCommand(versionCmd)
	mainCmd.AddCommand(serverCmd)

	flags := mainCmd.Flags()

	flags.String("db-host", "localhost", "DB Host")
	viper.BindPFlag("db_host", flags.Lookup("db-host"))

}

func main() {
	mainCmd.Execute()
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run Server",
	Run: func(cmd *cobra.Command, args []string) {

		// EXPORT USERLIST_DB_HOST=hop
		viper.SetEnvPrefix("userlist")
		viper.AutomaticEnv()

		log.SetLevel(log.DebugLevel)

		log.Info("Starting...")

		c := &config.ServerConfig{
			DBHost:     viper.GetString("db_host"),
			DBUser:     "postgres",
			DBPassword: "postgres",
			DBName:     "postgres",
			DBSSLMode:  "disable",
			HTTPPort:   8080,
			AdminList:  []string{"fsamin"},
		}

		models.DBConfig(c)

		e := server.New(c)
		e.Start()
	},
}
