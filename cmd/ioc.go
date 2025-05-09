/*
Copyright © 2025 Mason Clemons clemiondev@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// iocCmd represents the ioc command
var iocCmd = &cobra.Command{
	Use:   "ioc",
	Short: "Extracts indicators of compromise from a given file or directory (only text based files for now)",
	Long: `A utility meant to extract indicators of compromise from a given file or directory. Such as source code files or .txt, .log, .csv, files etc. For example:
	secutils ioc /path/to/file.eml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ioc called")
	},
}

func init() {
	rootCmd.AddCommand(iocCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// iocCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// iocCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
