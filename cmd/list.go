/*
Copyright © 2024 Harald Müller <harald.mueller@evosoft.com>

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
	"github.com/innomotics/tmtd/pkg/process"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "print a list of lsonld files with its title",
	Run: func(cmd *cobra.Command, args []string) {
		process.List(cmd.Flag("searchPath").Value.String())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("searchPath", "s", "", "list of directories for source files")
}
