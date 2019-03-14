package cmd

import (
	"github.com/betterde/ects/config"
	"github.com/spf13/cobra"
	"log"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Run a master node service",
	Long: "Run a master node service on this server",
	Run: func(cmd *cobra.Command, args []string) {
		exist, permission, err := config.CheckConfigFile("/etc/ects/ects.conf")
		log.Println(exist, permission, err)
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
