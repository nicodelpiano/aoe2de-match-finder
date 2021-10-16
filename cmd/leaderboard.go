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

// leaderboardCmd represents the leaderboard command
var leaderboardCmd = &cobra.Command{
	Use:   "leaderboard",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		top, err := cmd.Flags().GetInt("top")
		if err != nil {
			log.Fatalf("%v", err)
		}
		leaderboard, err := api.FetchLeaderboard(top)
		if err != nil {
			log.Fatalf("%v", err)
		}
		for _, topPlayer := range leaderboard.Leaderboard {
			log.Printf("%s\n\n", topPlayer.Info())
			// matches, err := api.FetchMatches(topPlayer.ProfileID)
			// if err != nil {
			// 	log.Printf("%v", err)
			// 	continue
			// }
			// matchesWonWithFranks := matches.Filter(api.PlayerWonWithFranks(topPlayer.Name))
			// for _, m := range matchesWonWithFranks {
			// 	log.Printf("Match found:")
			// 	log.Printf("%s", m.MatchID)
			// 	if err := m.Download(); err != nil {
			// 		log.Printf("Error downloading rec: %v", err)
			// 	}
			// }
		}
	},
}

func init() {
	rootCmd.AddCommand(leaderboardCmd)
	leaderboardCmd.Flags().IntP("top", "t", 100, "Top number")
}
