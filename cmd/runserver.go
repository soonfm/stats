package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Runserver CLI Flags
var (
	bind string
	port string
)

var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "Run the Stats Web Service.",
	Run:   func(c *cobra.Command, args []string) {},
}

// Sets default values for the runserver configuration
func runserverDefaults() {
	viper.SetDefault("server.bind", "0.0.0.0") // bind to all interfaces
	viper.SetDefault("server.port", "5000")    // port to serve on
}

func init() {
	// CLI Flags
	runserverCmd.Flags().StringVarP(&bind, "bind", "b", "", "IP to bind too (0.0.0.0)")
	runserverCmd.Flags().StringVarP(&port, "port", "p", "", "Port to listen on (5000)")

	// Set defaults
	runserverDefaults()

	// Bind Flags to Viper Configuration
	viper.BindPFlag("server.bind", runserverCmd.Flags().Lookup("bind"))
	viper.BindPFlag("server.port", runserverCmd.Flags().Lookup("port"))
}
