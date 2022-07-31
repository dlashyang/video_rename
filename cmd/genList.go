/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/dlashyang/video_rename/util"

	"github.com/spf13/cobra"
)

// genListCmd represents the genList command
var genListCmd = &cobra.Command{
	Use:   "genList path_to_videos",
	Short: "To generate the candidate file list, which are to be renamed.",
	Long: `To generate the candidate file list, which are to be renamed.

if no name for the candidate list is given, default name "mylist" 
would be used.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file_name, _ := cmd.Flags().GetString("list")
		fmt.Println("genList called")
		fmt.Println("candidate list is: ", file_name)
		util.Gen_candidate_list(string(args[0]), file_name)
	},
}

func init() {
	rootCmd.AddCommand(genListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	genListCmd.Flags().StringP("list", "l", "./mylist", "specify the list name, default is mylist")
}
