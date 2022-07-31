/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "video_rename",
	Short: "A tool to rename video files based on the capture date/time.",
	Long: `This tool can rename video files under a specific path(typically .mp4).

There are two steps are needed for the complete operation. 
First step is to use command "genList" to generate the candidate file list. 
Second step is to use command "rename" to actually rename video files based
on the generated list.
Between the two steps, the candidate file list(JSON format) can be modified 
manually.

Here is an example:
video_rename genList path
video_rename rename mylist.json`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) { //fmt.Println("main app called") },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.video_rename.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("test", "t", false, "test application")
}
