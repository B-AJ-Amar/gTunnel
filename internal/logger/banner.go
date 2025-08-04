package logger

import (
	"fmt"
	"strings"

	version "github.com/B-AJ-Amar/gTunnel/internal/pkg"
	"github.com/fatih/color"
)

func ShowBanner(mode string) {
	versionInfo := version.Get()

	cyan := color.New(color.FgCyan, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	green := color.New(color.FgGreen)
	blue := color.New(color.FgBlue)
	magenta := color.New(color.FgMagenta)

	banner := `
 ██████╗ ████████╗██╗   ██╗███╗   ██╗███╗   ██╗███████╗██╗     
██╔════╝ ╚══██╔══╝██║   ██║████╗  ██║████╗  ██║██╔════╝██║     
██║  ███╗   ██║   ██║   ██║██╔██╗ ██║██╔██╗ ██║█████╗  ██║     
██║   ██║   ██║   ██║   ██║██║╚██╗██║██║╚██╗██║██╔══╝  ██║     
╚██████╔╝   ██║   ╚██████╔╝██║ ╚████║██║ ╚████║███████╗███████╗
 ╚═════╝    ╚═╝    ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═══╝╚══════╝╚══════╝
`

	fmt.Println()
	cyan.Print(banner)

	if mode == "server" {
		yellow.Println("                      🚀 gTunnel Server 🚀")
	} else if mode == "client" {
		yellow.Println("                      🔗 gTunnel Client 🔗")
	}

	fmt.Println()
	green.Printf("  Version: %s", versionInfo.Version)
	if versionInfo.GitCommit != "unknown" && len(versionInfo.GitCommit) > 7 {
		green.Printf(" (%s)", versionInfo.GitCommit[:7])
	}
	fmt.Println()

	if versionInfo.BuildDate != "unknown" {
		blue.Printf("  Built: %s\n", versionInfo.BuildDate)
	}

	if versionInfo.GoVersion != "unknown" {
		blue.Printf("  Go: %s\n", versionInfo.GoVersion)
	}

	fmt.Println()
	magenta.Println(strings.Repeat("━", 70))
	fmt.Println()
}

func ShowSimpleBanner(mode string) {
	versionInfo := version.Get()

	fmt.Println()
	fmt.Println("================================================================")
	fmt.Println("                          gTunnel")
	if mode == "server" {
		fmt.Println("                        Server Mode")
	} else if mode == "client" {
		fmt.Println("                        Client Mode")
	}
	fmt.Println("================================================================")
	fmt.Printf("Version: %s", versionInfo.Version)
	if versionInfo.GitCommit != "unknown" && len(versionInfo.GitCommit) > 7 {
		fmt.Printf(" (%s)", versionInfo.GitCommit[:7])
	}
	fmt.Println()
	if versionInfo.BuildDate != "unknown" {
		fmt.Printf("Built: %s\n", versionInfo.BuildDate)
	}
	if versionInfo.GoVersion != "unknown" {
		fmt.Printf("Go: %s\n", versionInfo.GoVersion)
	}
	fmt.Println("================================================================")
	fmt.Println()
}
