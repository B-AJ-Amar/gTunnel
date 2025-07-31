package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display server status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Server is running. (Mock status)") // TODO: Real status check
	},
}
