package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Global Flags
var (
	configPath string
	logLevel   string
)

const DESCRIPTION string = "SFM 2.0 Stats Service\n\n" +
	"Collects usage metrics and provides an HTTP REST API."

// Root CLI Entrypoint command
var RootCmd = &cobra.Command{
	Use:   "statsrv",
	Short: "SFM Stats Service",
	Long:  DESCRIPTION,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := configure(); err != nil {
			logrus.Fatal("Failed to read configuration:", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Persistent Flags
	RootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Absolute Config File Path")
	RootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "", "Log Level (debug, info, warn, error)")

	// Set default configuration values
	setDefaults()

	// Bind Flags to Viper Configuration
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("logger.level", RootCmd.PersistentFlags().Lookup("log-level"))

	// Add Sub Commands
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(runserverCmd)
}

// Set default configuration values
func setDefaults() {
	// Logger Defaults
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.format", "text")
	viper.SetDefault("logger.logfile", "")
}

// Applies congiguration from a file, either passed directly in from the shell
// or from predefined lookup paths.
func configure() error {
	// Attempt to first read config from specific path if set
	if viper.IsSet("config") {
		viper.SetConfigFile(viper.GetString("config"))
		if err := viper.MergeInConfig(); err != nil {
			return err
		}
	} else { // No specific config path, try common locations
		// The name of the config should be config.json, config.yml etc
		viper.SetConfigName("config")
		// First try and load configuration from these paths
		viper.AddConfigPath("/etc/sfm/stats")
		viper.AddConfigPath("$HOME/.config/sfm/stats")
		// Attempt to read from these configs
		viper.MergeInConfig()
	}

	return nil
}

// Wrapper around RootCmd.Execute to set version from build argument
func Execute(version string) error {
	VERSION = version // see version.go
	return RootCmd.Execute()
}
