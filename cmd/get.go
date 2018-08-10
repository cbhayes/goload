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
	"net/http"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var template *string
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

		for i := 0; i < *numRequests; i++ {
			status := make(chan int)
			fmt.Printf("Starting request %d :\n", i)
			for j := 0; j < *numClients; j++ {
				go sendRequest(urlString, nil, status)
			}

			for j := 0; j < *numClients; j++ {
				fmt.Printf("Client %d Status %d\n", j, <-status)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	template = getCmd.Flags().StringP("template", "t", "", "Path to template file")
	numRequests = getCmd.Flags().IntP("numRequests", "n",10, "Number of requests to send")
	numClients = getCmd.Flags().IntP("numClients",  "c",1, "Number of clients to send from")
}

func sendRequest(url string, template []byte, c chan int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request Failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read body")
	}
	if body != nil {
		fmt.Printf("%s", body)
	}
	c <- resp.StatusCode

	//"{\"firstname\":\"[firstname]\",\"lastname\":\"[lastname]\",\"age\":[number]}"
}
