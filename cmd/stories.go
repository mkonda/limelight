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

const api string = "https://api.app.shortcut.com/api/v3/search/stories"

// storiesCmd represents the github command
var storiesCmd = &cobra.Command{
	Use:   "stories",
	Short: "Get stories from Shortcut",
	Long:  `Pull story information out of Shortcut.`,
	Run: func(cmd *cobra.Command, args []string) {
		shortcuttoken := viper.GetString("shortcut-token")
		q := viper.GetString("stories-query")
		query := Query{25, q}
		jsonQuery, _ := json.Marshal(query)
		client := &http.Client{}
		req, _ := http.NewRequest("GET", api, bytes.NewBuffer(jsonQuery))
		req.Header.Set("Shortcut-Token", shortcuttoken)
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)

		// Convert response body to string
		//bodyString := string(bodyBytes)
		//fmt.Println("API Response as String:\n" + bodyString)

		// Convert response body to Response struct
		var responseStruct StoryResponse
		json.Unmarshal(bodyBytes, &responseStruct)
		stories := responseStruct.Data
		for i := 0; i < len(stories); i++ {
			fmt.Printf("Project: %v\tEpic: %v\n\tStory: %v\tDue: %v\tID: %v\tState: %v\tLabels: %v\n\tURL: %v\n", getProject(stories[i].ProjectID).Name,
				getEpic(stories[i].EpicID).Name, stories[i].Name, stories[i].Deadline,
				stories[i].ID, stories[i].WorkflowStateID, stories[i].Labels, stories[i].URL)
		}
	},
}

func init() {
	shortcutCmd.AddCommand(storiesCmd)
	storiesCmd.Flags().String("stories-query", "", "The story query to run.")
	viper.BindPFlag("stories-query", storiesCmd.Flags().Lookup("stories-query"))
}
