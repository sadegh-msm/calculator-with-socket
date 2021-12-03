package cmd

import "github.com/spf13/cobra"

// rootCommand is an internal command
var rootCommand = &cobra.Command{
	Use:   "root",
	Short: "short description",
	Long:  "long description",
}

func Execute() {
	cobra.CheckErr(rootCommand.Execute())
}
