// Copyright © 2019-2021 Matt Konda mkonda@jemurai.com
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

	"github.com/spf13/cobra"
)

var versionString = "0.0.2"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version of limelight",
	Long:  `Displays the version of limelight`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Limelight Version", versionString)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
