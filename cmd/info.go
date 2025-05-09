/*
Copyright © 2025 Mason Clemons clemiondev@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Analyzes the given file or directory and returns information about it",
	Long: `A utility meant to give information relevant information to cybersecurity practicioners about a given file or directory. For example:
	secutils info /path/to/file
	-or- 
	secutils info /path/to/directory
	(leaving the file or directory argument blank [secutils info]  will default to the current working directory secutils info is executed in)`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
