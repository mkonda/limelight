// Package cmd

package shortcut

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

)

const workflowsapi string = "https://api.app.shortcut.com/api/v3/workflows"

// statesCmd represents the github command
var statesCmd = &cobra.Command{
	Use:   "states",
	Short: "Get the workflow states from Shortcut",
	Long:  `Pull workflow status out of Shortcut.`,
	Run: func(cmd *cobra.Command, args []string) {
		shortcuttoken := viper.GetString("shortcut-token")
		client := &http.Client{}
		req, _ := http.NewRequest("GET", workflowsapi, nil)
		req.Header.Set("Shortcut-Token", shortcuttoken)
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)

		// Convert response body to Response struct
		var workflows []Workflow
		json.Unmarshal(bodyBytes, &workflows)
		for i := 0; i < len(workflows); i++ {
			workflow := workflows[i]
			fmt.Printf("\nWorkflow: %v\n", workflow.Name)
			states := workflow.States
			for j := 0; j < len(states); j++ {
				fmt.Printf("Workflow State: \tName: %v\tID: %v\n", states[j].Name, states[j].ID)
			}
		}
	},
}

func init() {
	shortcutCmd.AddCommand(statesCmd)
}
