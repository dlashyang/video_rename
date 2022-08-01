/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/dlashyang/video_rename/util"
	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename video files based on the generated list",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rename called")
		flagDryRun, _ := cmd.Flags().GetBool("dry-run")
		fmt.Println("dry-run? ", flagDryRun)
		listName, _ := cmd.Flags().GetString("list")
		fmt.Println("list file: ", listName)
		util.RenamebyList(listName, flagDryRun)
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	renameCmd.Flags().BoolP("dry-run", "n", false, "Print only, no real actions")
	renameCmd.Flags().StringP("list", "l", "./mylist", "specify the list name, default is mylist")
}
