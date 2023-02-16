// Package cmd

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const nextapi string = "https://api.app.shortcut.com/api/v3/search/stories"

// storiesCmd represents the github command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Get the next story to work on from Shortcut",
	Long:  `Pull story information out of Shortcut for the next thing I'm supposed to do`,
	Run: func(cmd *cobra.Command, args []string) {
		shortcuttoken := viper.GetString("shortcut-token")
		q := viper.GetString("next-query")
		query := Query{2, q}
		jsonQuery, _ := json.Marshal(query)
		client := &http.Client{}
		req, _ := http.NewRequest("GET", nextapi, bytes.NewBuffer(jsonQuery))
		req.Header.Set("Shortcut-Token", shortcuttoken)
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)

		// Convert response body to Response struct
		var responseStruct StoryResponse
		json.Unmarshal(bodyBytes, &responseStruct)
		stories := responseStruct.Data
		for i := 0; i < 1; i++ {  // A dumb hack ... but it works.  :)
			if (len(stories) < 1) {
				fmt.Printf("No Stories Found\n")
			}else {
				fmt.Printf("Project: %v\n\tEpic: %v\n\tStory: %v\n\tDue: %v\n\tID: %v\n\tState: %v\n\tLabels: %v\n\tURL: %v\n\n", getProject(stories[i].ProjectID).Name,
				getEpic(stories[i].EpicID).Name, stories[i].Name, stories[i].Deadline,
				stories[i].ID, stories[i].WorkflowStateID, stories[i].Labels, stories[i].URL)
			}
			}
	},
}

func init() {
	shortcutCmd.AddCommand(nextCmd)
	nextCmd.Flags().String("next-query", "", "The next story query to run.")
	viper.BindPFlag("next-query", nextCmd.Flags().Lookup("next-query"))
}
