// Copyright © 2018 Conor Hayes <hayesconorb@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/conchuirh/goload/measure"
	"github.com/conchuirh/goload/template"
	"github.com/spf13/cobra"
)

var templateString *string
var numRequests *int
var numClients *int

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Send GET requests",
	Long: `Send GET requests to destination URL, using body from template if provided
For example:
goload get http://example.com/path -n 2000 -c 10
will send 2000 GET requests each for 10 clients to http://example.com/path`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		urlString := args[0]

		fmt.Printf("Starting %d requests for %d clients to %s \n", *numRequests, *numClients, urlString)

		c := make(chan measure.Measure)
		for client := 0; client < *numClients; client++ {
			go createClient(client, *numRequests, urlString, c)
		}

		measures := make([]measure.Measure, *numRequests**numClients)
		for i := *numRequests * *numClients; i > 0; i-- {
			measures[i-1] = <-c
		}

		measure.Stats(&measures)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	templateString = getCmd.Flags().StringP("template", "t", "", "Path to template file")
	numRequests = getCmd.Flags().IntP("numRequests", "n", 10, "Number of requests to send")
	numClients = getCmd.Flags().IntP("numClients", "c", 1, "Number of clients to send from")
}

func createClient(client int, numReq int, url string, c chan measure.Measure) {
	for req := 0; req < numReq; req++ {
		m := sendRequest(url, nil)
		fmt.Printf("Client %d Request #%d Duration %s\n", client, req, m.Elapsed.String())
		c <- m
	}
	return
}

func sendRequest(url string, templateBytes []byte) measure.Measure {
	start := time.Now()
	resp, err := http.Get(template.Build(url))
	if err != nil {
		fmt.Println("Request Failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read body")
	}
	end := time.Now()
	return measure.Create(start, end, resp.StatusCode, body)
}
