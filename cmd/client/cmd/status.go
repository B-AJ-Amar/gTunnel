package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display client connection status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Client connection status: Not connected (Mock status)") // TODO: Real status check
	},
}
