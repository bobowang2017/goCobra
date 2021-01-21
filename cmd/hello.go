package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

func init() {
	rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello World",
	Long:  "Hello The All World",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello : " + strconv.Itoa(i))
			time.Sleep(10000 * time.Microsecond)
		}
	},
}
