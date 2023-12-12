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
	"io"
	"os"
	"strings"
	"vimbin/internal/config"
	"vimbin/internal/utils"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

// fetchCmd represents the 'fetch' command for retrieving the latest data from the vimbin server.
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches the latest data from the vimbin server",
	Long: `The 'fetch' command retrieves the latest content from the vimbin server specified by the provided URL.
It makes a GET request to the server and prints the response body to the console.

Example:
  vimbin fetch --url http://example.com`,
	Run: func(cmd *cobra.Command, args []string) {
		url := strings.TrimSuffix(config.App.Server.Api.Address, "/")

		if url == "" {
			log.Error().Msg("URL is empty")
			os.Exit(1)
		}
		url += "/fetch"

		httpClient := utils.CreateHTTPClient(config.App.Server.Api.SkipInsecureVerify)

		// Make a GET request to the vimbin server
		response, err := httpClient.Get(url)
		if err != nil {
			log.Error().Msgf("Error making GET request: %s", err)
			return
		}
		defer response.Body.Close()

		// Read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Error().Msgf("Error reading response body: %s", err)
			return
		}

		// Print the content to the console
		fmt.Println(string(body))
	},
}

func init() {
	// Add 'fetchCmd' to the root command
	rootCmd.AddCommand(fetchCmd)

	// Define command-line flags for 'fetchCmd'
	fetchCmd.PersistentFlags().StringVarP(&config.App.Server.Api.Address, "url", "u", "", "The URL of the vimbin server")
	fetchCmd.PersistentFlags().BoolVarP(&config.App.Server.Api.SkipInsecureVerify, "insecure-skip-verify", "i", false, "Skip TLS certificate verification")
}
