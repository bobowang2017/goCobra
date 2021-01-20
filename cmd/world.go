package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(worldCmd)
}

var worldCmd = &cobra.Command{
	Use:   "world",
	Short: "Hello World",
	Long:  "Hello The All World",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("image args are : " + strings.Join(args, " "))
	},
}
