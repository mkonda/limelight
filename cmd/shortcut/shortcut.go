// Package cmd

package shortcut

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mkonda/limelight/cmd"
)

// Query to send to shortcut
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

// Story Change struct
type StoryChange struct {
	Label    CreateLabelParams `json:"labels"`
	NewState int               `json:"workflow_state_id"`
}

// Story structure
type Story struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Deadline        string  `json:"deadline"`
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

// Epic struct
// https://api.app.shortcut.com/api/v3/epics/{epic-public-id}
type Epic struct {
	Name string `json:"name"`
}

// Project struct
// https://api.app.shortcut.com/api/v3/projects/{project-public-id}
type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Create Label Parameters
type CreateLabelParams struct {
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

// These are probably specific to our project.  We're not using them yet, but
// you might want to look at the states command to dump them.

// Kanban
const UnstartedKanbanState int = 500000027
const StartedKanbanState int = 500000026
const DoneKanbanState int = 500000028

// Engineering
const UnscheduledEngineeringState int = 500000008
const ReadyForDevelopmentState int = 500000007
const ReadyForReviewEngineeringState int = 500000010
const ReadyForDeployEngineeringState int = 500000009
const CompletedEngineeringState int = 500000011

// Workflow
type Workflow struct {
	ID     int             `json:"id"`
	Name   string          `json:"name"`
	States []WorkflowState `json:"states"`
}

// WorkflowState
type WorkflowState struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// storiesCmd represents the github command
var shortcutCmd = &cobra.Command{
	Use:   "shortcut",
	Short: "Connect to shortcut and do something",
	Long:  `Connect to shortcut and do something.  (See subcommands)`,
}

func init() {
	cmd.RootCmd.AddCommand(shortcutCmd)
	shortcutCmd.Flags().String("shortcut-token", "", "The token to connect with.")
	viper.BindPFlag("shortcut-token", shortcutCmd.Flags().Lookup("shortcut-token"))
}

func getEpic(ID int) Epic {
	shortcuttoken := viper.GetString("shortcut-token")
	epicAPI := "https://api.app.shortcut.com/api/v3/epics/" + strconv.Itoa(ID)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", epicAPI, nil)
	req.Header.Set("Shortcut-Token", shortcuttoken)
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
	shortcuttoken := viper.GetString("shortcut-token")
	projectAPI := "https://api.app.shortcut.com/api/v3/projects/" + strconv.Itoa(ID)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", projectAPI, nil)
	req.Header.Set("Shortcut-Token", shortcuttoken)
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
