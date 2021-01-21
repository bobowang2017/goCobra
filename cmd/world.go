package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rootCmd.AddCommand(worldCmd)
}

var worldCmd = &cobra.Command{
	Use:   "world",
	Short: "Hello World",
	Long:  "Hello The All World",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < 10000; i++ {
			fmt.Println("World : " + strconv.Itoa(i))
		}
	},
}
