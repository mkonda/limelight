// Package cmd

package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Query to send to clubhouse
type Query struct {
	PageSize int    `json:"page_size"`
	Query    string `json:"query"`
}

// StoryResponse structure
type StoryResponse struct {
	Next string  `json:"next"`
	Data []Story `json:"data"`
}

// Label structure
type Label struct {
	Name string `json:"name"`
}

// Story structure
type Story struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Deadline        string  `json:"deadline"`
	ID              int     `json:"id"`
	Labels          []Label `json:"labels`
	URL             string  `json:"app_url"`
	ProjectID       int     `json:"project_id"`
	EpicID          int     `json:"epic_id"`
	WorkflowStateID int     `json:"workflow_state_id"` // This is kind of annoyingly useless.
	// These are not in JSON
	ProjectName string
	EpicName    string
	State       string
}

// Project struct
// https://api.clubhouse.io/api/v3/projects/{project-public-id}
type Project struct {
	Name string `json:"name"`
}

// Epic struct
// https://api.clubhouse.io/api/v3/epics/{epic-public-id}
type Epic struct {
	Name string `json:"name"`
}

// Epic struct
// https://api.clubhouse.io/api/v3/epics/{epic-public-id}
type State struct {
	Name string `json:"name"`
}

// storiesCmd represents the github command
var clubhouseCmd = &cobra.Command{
	Use:   "clubhouse",
	Short: "Connect to clubhouse and do something",
	Long:  `Connect to clubhouse and do something.  (See subcommands)`,
}

func init() {
	rootCmd.AddCommand(clubhouseCmd)

	clubhouseCmd.Flags().String("clubhouse-token", "", "The token to connect with.")
	viper.BindPFlag("clubhouse-token", clubhouseCmd.Flags().Lookup("clubhouse-token"))

}

func getEpic(ID int) Epic {
	clubhousetoken := viper.GetString("clubhouse-token")
	epicAPI := "https://api.clubhouse.io/api/v3/epics/" + strconv.Itoa(ID)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", epicAPI, nil)
	req.Header.Set("Clubhouse-Token", clubhousetoken)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	// Convert response body to string
	//	bodyString := string(bodyBytes)
	//	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Response struct
	var responseStruct Epic
	json.Unmarshal(bodyBytes, &responseStruct)
	return responseStruct
}

func getProject(ID int) Project {
	clubhousetoken := viper.GetString("clubhouse-token")
	projectAPI := "https://api.clubhouse.io/api/v3/projects/" + strconv.Itoa(ID)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", projectAPI, nil)
	req.Header.Set("Clubhouse-Token", clubhousetoken)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	// Convert response body to string
	//	bodyString := string(bodyBytes)
	//	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Response struct
	var responseStruct Project
	json.Unmarshal(bodyBytes, &responseStruct)
	return responseStruct
}
