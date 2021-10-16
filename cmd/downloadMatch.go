/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/nicodelpiano/aoe2de-match-finder/api"
	"github.com/spf13/cobra"
)

// downloadMatchCmd represents the downloadMatch command
var downloadMatchCmd = &cobra.Command{
	Use:   "downloadMatch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		matchID, err := cmd.Flags().GetInt("match-id")
		log.Printf("id: %d", matchID)
		if err != nil {
			log.Fatalf("a %v", err)
		}
		m, err := api.FetchMatch(matchID)
		if err != nil {
			log.Fatalf("b %v", err)
		}
		err = m.Download()
		if err != nil {
			log.Fatalf("c %v", err)
		}
		log.Printf("Successfully downloaded match %d", matchID)
	},
}

func init() {
	rootCmd.AddCommand(downloadMatchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadMatchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadMatchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	downloadMatchCmd.Flags().IntP("match-id", "m", 0, "Match ID to search")
}
