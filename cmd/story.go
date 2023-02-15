// Package cmd

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const base_story_api string = "https://api.app.shortcut.com/api/v3/stories/"

// storiesCmd represents the github command
var storyCmd = &cobra.Command{
	Use:   "story",
	Short: "Update story in Clubhouse",
	Long:  `Update story in Clubhouse.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		clubhousetoken := viper.GetString("clubhouse-token")
		storyID := args[0]
		putAPI := base_story_api + storyID
		storyState := viper.GetString("story-state")
		storyLabel := viper.GetString("story-label")
		var change string
		if storyState != "" && storyLabel == "" {
			change = fmt.Sprintf("{ \"workflow_state_id\": %v }", storyState)
		} else if storyState != "" && storyLabel != "" {
			change = fmt.Sprintf("{ \"workflow_state_id\": %v, labels\": [{ \"color\": \"#cc5856\", \"description\": \"%v\", \"external_id\": null, \"name\": \"%v\" }] }", storyState, storyLabel, storyLabel)
		} else if storyState == "" && storyLabel != "" {
			change = fmt.Sprintf("{ \"labels\": [{ \"color\": \"#cc5856\", \"description\": \"%v\", \"name\": \"%v\" }] }", storyLabel, storyLabel)
		} else {
			change = "{}"
		}

		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodPut, putAPI, bytes.NewBuffer([]byte(change)))
		req.Header.Set("Clubhouse-Token", clubhousetoken)
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)

		bodyString := string(bodyBytes)
		fmt.Println("API Response as String:\n" + bodyString)

	},
}

func init() {
	shortcutCmd.AddCommand(storyCmd)
	storyCmd.Flags().String("story-state", "", "The state to set the story to.")
	viper.BindPFlag("story-state", storyCmd.Flags().Lookup("story-state"))
	storyCmd.Flags().String("story-label", "", "The label to add to the story.")
	viper.BindPFlag("story-label", storyCmd.Flags().Lookup("story-label"))
}

/*
func getStoryState(state string) (int, error) {
	newState, err := strconv.Atoi(state)
	if err != nil {
		fmt.Errorf("Error parsing target state %v", state)
		return 0, errors.New("Invalid state for target new story state")
	}
	// TODO:  Add checking for actual states...
	return newState, nil
}

func getStoryLabel(label string) (CreateLabelParams, error) {
	color := "#49a940" // default is green
	if label == "" {
		return CreateLabelParams{"Test", color, "Test"}, errors.New("No Label Defined")
	}
	if label == "PriorityA" {
		color = "#cc5856" // Set PriorityA to Red ...
	}
	return CreateLabelParams{label, color, label}, nil
}
*/
