package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var outEnvCmd = &cobra.Command{
	Use:   "outenv",
	Short: "Output all enviroment variables to std",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Out....")
	},
}
