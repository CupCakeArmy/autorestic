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
	"github.com/cupcakearmy/autorestic/internal"
	"github.com/spf13/cobra"
)

// forgetCmd represents the forget command
var forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Forget and optionally prune snapshots according the specified policies",
	Run: func(cmd *cobra.Command, args []string) {
		config := internal.GetConfig()
		if err := config.CheckConfig(); err != nil {
			panic(err)
		}
		{
			selected, err := internal.GetAllOrSelected(cmd, false)
			cobra.CheckErr(err)
			prune, _ := cmd.Flags().GetBool("prune")
			for _, name := range selected {
				location := config.Locations[name]
				err := location.Forget(prune)
				cobra.CheckErr(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(forgetCmd)
	internal.AddFlagsToCommand(forgetCmd, false)
	forgetCmd.Flags().Bool("prune", false, "Also prune repository")
}