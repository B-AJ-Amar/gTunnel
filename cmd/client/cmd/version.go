package cmd

import (
	"encoding/json"
	"fmt"

	version "github.com/B-AJ-Amar/gTunnel/internal/pkg"
	"github.com/spf13/cobra"
)

var (
	versionOutputFormat string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `Display detailed version information including build details.`,
	RunE:  runVersion,
}

func runVersion(cmd *cobra.Command, args []string) error {
	versionInfo := version.Get()

	switch versionOutputFormat {
	case "json":
		jsonOutput, err := json.MarshalIndent(versionInfo, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal version info: %w", err)
		}
		fmt.Println(string(jsonOutput))
	case "short":
		fmt.Println(versionInfo.Version)
	default:
		fmt.Printf("gTunnel Client\n")
		fmt.Printf("Version:    %s\n", versionInfo.Version)
		fmt.Printf("Git Commit: %s\n", versionInfo.GitCommit)
		fmt.Printf("Build Date: %s\n", versionInfo.BuildDate)
		fmt.Printf("Go Version: %s\n", versionInfo.GoVersion)
	}

	return nil
}

func init() {
	versionCmd.Flags().StringVarP(&versionOutputFormat, "output", "o", "default",
		"Output format. One of: default|json|short")
}
