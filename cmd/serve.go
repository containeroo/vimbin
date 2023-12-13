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
	"strings"
	"text/template"
	"vimbin/internal/config"
	"vimbin/internal/handlers"
	"vimbin/internal/server"
	"vimbin/internal/utils"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

// supportedThemes is a list of themes supported by the serve command.
var supportedThemes = []string{"auto", "light", "dark"}

// serveCmd represents the serve command.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long: `Start a local web server featuring a textarea for creating, managing, and refining content.
  All changes made in the textarea are persistently stored to a file, and users can navigate and
  manipulate text using familiar Vim motions for an enhanced editing experience.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the specified theme is supported
		if !utils.IsInList(config.App.Server.Web.Theme, supportedThemes) {
			fmt.Printf("Unsupported output format: %s. Supported formats are: %s\n", config.App.Server.Web.Theme, strings.Join(supportedThemes, ", "))
			_ = cmd.Help() // Makes the linter happy
			os.Exit(1)
		}

		// Read the HTML template file
		htmlTemplate, err := template.ParseFS(server.StaticFS, "web/templates/index.html")
		if err != nil {
			log.Fatal().Err(err)
		}

		config.App.HtmlTemplate = htmlTemplate

		// Parse the configuration
		if err := config.App.Parse(); err != nil {
			log.Fatal().Err(err)
		}

		// Collect handlers and start the server
		handlers.Collect()
		server.Run(config.App.Server.Web.Address)
	},
}

func init() {
	// Add the serve command to the root command
	rootCmd.AddCommand(serveCmd)

	// Define command-line flags for the serve command
	serveCmd.PersistentFlags().StringVarP(&config.App.Server.Web.Address, "listen-address", "a", ":8080", "The address to listen on for HTTP requests.")

	serveCmd.PersistentFlags().StringVarP(&config.App.Server.Web.Theme, "theme", "t", "auto", fmt.Sprintf("The theme to use. Can be %s.", strings.Join(supportedThemes, "|")))
	serveCmd.RegisterFlagCompletionFunc("theme", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return supportedThemes, cobra.ShellCompDirectiveDefault
	})

	serveCmd.PersistentFlags().StringVarP(&config.App.Storage.Directory, "directory", "d", "$(pwd)", "The path to the storage directory. Defaults to the current working directory.")
	serveCmd.PersistentFlags().StringVarP(&config.App.Storage.Name, "name", "n", ".vimbin", "The name of the file to save.")
}
