/*
Copyright © 2023 containeroo hello©containeroo.ch

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"time"
	"vimbin/internal/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

const (
	version = "v0.0.11"
)

var (
	cfgFile      string
	debug        bool
	trace        bool
	printVersion bool
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "vimbin",
	Short: "vimbin - a pastebin with vim motion",
	Long: `vimbin is a powerful pastebin tool that seamlessly integrates the efficiency of Vim motions
	with the simplicity of a pastebin service. It offers two main commands:

- server: Start a local web server featuring a textarea for creating, managing, and refining content.
  All changes made in the textarea are persistently stored to a file, and users can navigate and
  manipulate text using familiar Vim motions for an enhanced editing experience.

- push: Quickly send text to the vimbin server from the command line. This allows for easy integration
  with other tools and scripts, streamlining the process of sharing content through vimbin.
  `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// PersistentPreRun is executed before any subcommand is executed.

		// Check if the version flag is set and print the version
		if printVersion {
			fmt.Println("vimbin version:", version)
			os.Exit(0)
		}

		// Configure zerolog
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		log.Logger = zerolog.New(output).With().Timestamp().Logger()

		// Configure log levels based on debug and trace flags
		if debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			log.Debug().Msgf("Verbose output enabled")
		} else if trace {
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
			log.Debug().Msgf("Trace output enabled")
		}
		if token := cmd.Flag("token").Value.String(); token != "" {
			config.App.Server.Api.Token.Set(token)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Run is executed if no subcommand is specified.

		// Check if the version flag is set and print the version
		if printVersion {
			fmt.Println("vimbin version:", version)
			os.Exit(0)
		} else {
			// Display the root command's help message
			_ = cmd.Help() // Make the linter happy
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once for the rootCmd.
func Execute() {
	// Execute the root command and handle any errors
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Initialize the configuration
	cobra.OnInitialize(initConfig)

	// Define command-line flags for the root command
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Path to the configuration file.")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "", false, "Activates debug output for detailed logging.")
	rootCmd.PersistentFlags().StringP("token", "t", "", "Token to use for authentication. If not set, a random token will be generated.")
	rootCmd.PersistentFlags().BoolVarP(&trace, "trace", "", false, "Enables trace mode. This will show the content in the logs!")
	rootCmd.MarkFlagsMutuallyExclusive("debug", "trace") // Ensure that debug and trace flags are mutually exclusive

	rootCmd.PersistentFlags().BoolVarP(&printVersion, "version", "v", false, "Print version and exit.")
}

// initConfig reads the configuration from the specified file or environment variables.
func initConfig() {
	if cfgFile != "" {
		// Use the config file specified by the flag.
		viper.SetConfigFile(cfgFile)
		viper.AutomaticEnv() // Read in environment variables that match

		if err := config.App.Read(viper.ConfigFileUsed()); err != nil {
			log.Fatal().Msgf("Error reading config file: %v", err)
		}
	}
}
