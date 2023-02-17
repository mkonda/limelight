// Package cmd

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const projectsapi string = "https://api.app.shortcut.com/api/v3/projects"

// projectsCmd represents the github command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Get the projects  from Shortcut",
	Long:  `Pull project names and ids out of Shortcut.`,
	Run: func(cmd *cobra.Command, args []string) {
		shortcuttoken := viper.GetString("shortcut-token")
		client := &http.Client{}
		req, _ := http.NewRequest("GET", projectsapi, nil)
		req.Header.Set("Shortcut-Token", shortcuttoken)
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(res.Body)

		// Convert response body to Response struct
		var projects []Project
		json.Unmarshal(bodyBytes, &projects)
		for i := 0; i < len(projects); i++ {
			project := projects[i]
			fmt.Printf("Project Name: %s, Project ID: %v\n", project.Name, project.ID)
		}
	},
}

func init() {
	shortcutCmd.AddCommand(projectsCmd)
}
