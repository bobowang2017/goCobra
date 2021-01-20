package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello World",
	Long:  "Hello The All World",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("image args are : " + strings.Join(args, " "))
	},
}
