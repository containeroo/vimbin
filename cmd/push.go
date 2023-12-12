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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"vimbin/internal/config"
	"vimbin/internal/utils"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

var appendFlag bool

// pushCmd represents the 'push' command for sending data to the vimbin server.
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Pushes data to the vimbin server",
	Long: `Push sends data to the vimbin server, allowing you to store text content.
It supports two modes: 'save' and 'append'. In 'save' mode, the entire content is
replaced, while in 'append' mode, new content is added to the existing content.

Examples:
  - Save content:
    vimbin push "Your text content" --url http://example.com
  - Append content:
    vimbin push --append "Additional content" --url http://example.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if at least one character is provided
		if len(args) < 1 {
			log.Error().Msg("You must push at least one character.")
			os.Exit(1)
		}

		url := strings.TrimSuffix(config.App.Server.Api.Address, "/")
		if url == "" {
			log.Error().Msg("URL is empty")
			os.Exit(1)
		}

		// Concatenate input arguments into a single string
		body := strings.Join(args, "\n")

		// Build the URL based on the "append" flag
		if appendFlag {
			url += "/append"
			body = "\n" + body
		} else {
			url += "/save"
		}

		// Prepare content for the POST request
		content := map[string]string{"content": body}
		requestBody, err := json.Marshal(content)
		if err != nil {
			log.Error().Msgf("Error encoding JSON: %s", err)
			os.Exit(1)
		}

		httpClient := utils.CreateHTTPClient(config.App.Server.Api.SkipInsecureVerify)

		// Make the POST request to the vimbin server
		response, err := httpClient.Post(url, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Error().Msgf("Error making POST request: %s", err)
			os.Exit(1)
		}
		defer response.Body.Close()

		// Check for successful response
		if response.StatusCode != http.StatusOK {
			log.Error().Msgf("Unexpected status code %d", response.StatusCode)
			os.Exit(1)
		}

		// Read and print the response body
		var responseBodyBuffer bytes.Buffer
		_, err = io.Copy(&responseBodyBuffer, response.Body)
		if err != nil {
			log.Error().Msgf("Error reading response body: %s", err)
			os.Exit(1)
		}

		fmt.Println(responseBodyBuffer.String())
	},
}

func init() {
	// Add 'fetchCmd' to the root command
	rootCmd.AddCommand(pushCmd)

	// Define command-line flags for 'fetchCmd'
	pushCmd.PersistentFlags().StringVarP(&config.App.Server.Api.Address, "url", "u", "", "The URL of the vimbin server")
	pushCmd.PersistentFlags().BoolVarP(&config.App.Server.Api.SkipInsecureVerify, "insecure-skip-verify", "i", false, "Skip TLS certificate verification")
	pushCmd.PersistentFlags().BoolVarP(&appendFlag, "append", "a", false, "Append content to the existing content")
}
