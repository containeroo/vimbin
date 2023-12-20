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
	"net/http"
	"strings"
	"vimbin/internal/config"
	"vimbin/internal/utils"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// pullCmd represents the 'fetch' command for retrieving the latest data from the vimbin server.
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls the latest data from the vimbin server",
	Long: `The 'pull' command retrieves the latest content from the vimbin server specified by the provided URL.
It makes a GET request to the server and prints the response body to the console.

Example:
  vimbin pull --url http://example.com`,
	Run: func(cmd *cobra.Command, args []string) {
		url := strings.TrimSuffix(config.App.Server.Api.Address, "/")

		if url == "" {
			log.Fatal().Msg("URL is empty")
		}
		url += "/fetch"
		log.Debug().Msgf("URL: %s", url)

		apiToken := config.App.Server.Api.Token.Get()
		if apiToken == "" {
			log.Fatal().Msg("API token is empty")
		}

		httpClient := utils.CreateHTTPClient(config.App.Server.Api.SkipInsecureVerify)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal().Msgf("Error creating HTTP request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-API-Token", apiToken)

		// Make a GET request to the vimbin server
		response, err := httpClient.Do(req)
		if err != nil {
			log.Error().Msgf("Error making GET request: %s", err)
			return
		}
		defer response.Body.Close()

		// Check for successful response
		if response.StatusCode != http.StatusOK {
			log.Fatal().Msgf("Unexpected status code %d", response.StatusCode)
		}

		// Read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal().Msgf("Error reading response body: %s", err)
		}

		// Print the content to the console
		fmt.Println(string(body))
	},
}

func init() {
	// Add 'pullCmd' to the root command
	rootCmd.AddCommand(pullCmd)

	// Define command-line flags for 'pullCmd'
	pullCmd.PersistentFlags().StringVarP(&config.App.Server.Api.Address, "url", "u", "", "The URL of the vimbin server")
	pullCmd.PersistentFlags().BoolVarP(&config.App.Server.Api.SkipInsecureVerify, "insecure-skip-verify", "i", false, "Skip TLS certificate verification")
}
