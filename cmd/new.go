// Package cmd

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const new_story_api string = "https://api.app.shortcut.com/api/v3/stories"

// newStoryCmd represents the github command
var newStoryCmd = &cobra.Command{
	Use:   "new",
	Short: "New story in Shortcut",
	Long:  `New story in Shortcut.  Pass the project (int), a name and a description for the story.`,
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		shortcuttoken := viper.GetString("shortcut-token")
		storyProject,err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln(err)
			log.Fatalln("Project ID should be an integer")
		}
		storyName := args[1]
		storyDesk := args[2]
		var change string = fmt.Sprintf("{  \"name\": \"%s\", \"project_id\": %d, \"description\": \"%s\" }",  storyName, storyProject, storyDesk)
		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodPost, new_story_api, bytes.NewBuffer([]byte(change)))
		req.Header.Set("Shortcut-Token", shortcuttoken)
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)

		// Useful for debugging...
		//		bodyString := string(bodyBytes)
		//		fmt.Println("API Response as String:\n" + bodyString)

		var responseStory Story
		json.Unmarshal(bodyBytes, &responseStory)
		fmt.Printf("ID: %v \tURL: %v\n", responseStory.ID, responseStory.URL)
	},
}

func init() {
	shortcutCmd.AddCommand(newStoryCmd)
}
